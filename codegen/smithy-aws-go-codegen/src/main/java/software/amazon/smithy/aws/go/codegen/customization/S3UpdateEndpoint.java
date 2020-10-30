/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://aws.amazon.com/apache2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 *
 *
 */

package software.amazon.smithy.aws.go.codegen.customization;

import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;
import software.amazon.smithy.aws.traits.ServiceTrait;
import software.amazon.smithy.codegen.core.CodegenException;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.go.codegen.GoDelegator;
import software.amazon.smithy.go.codegen.GoSettings;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.SymbolUtils;
import software.amazon.smithy.go.codegen.integration.ConfigField;
import software.amazon.smithy.go.codegen.integration.GoIntegration;
import software.amazon.smithy.go.codegen.integration.MiddlewareRegistrar;
import software.amazon.smithy.go.codegen.integration.RuntimeClientPlugin;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.OperationShape;
import software.amazon.smithy.model.shapes.ServiceShape;
import software.amazon.smithy.model.shapes.ShapeId;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.utils.ListUtils;
import software.amazon.smithy.utils.SetUtils;

/**
 * S3UpdateEndpoint integration serves to apply customizations for S3 service,
 * and modifies the resolved endpoint based on S3 client config or input shape values.
 */
public class S3UpdateEndpoint implements GoIntegration {
    // options to be generated on Client's options type
    private static final String USE_PATH_STYLE_OPTION = "UsePathStyle";
    private static final String USE_ACCELERATE_OPTION = "UseAccelerate";
    private static final String USE_DUALSTACK_OPTION = "UseDualstack";
    private static final String USE_ARNREGION_OPTION = "UseARNRegion";

    // middleware addition constants
    private static final String UPDATE_ENDPOINT_ADDER = "addUpdateEndpointMiddleware";
    private static final String UPDATE_ENDPOINT_INTERNAL_ADDER = "UpdateEndpoint";

    // private function getter constant
    private static final String GET_BUCKET_FROM_INPUT = "getBucketFromInput";
    private static final String SUPPORT_ACCELERATE = "supportAccelerate";

    // list of operations that do not support accelerate
    private static final Set<String> NOT_SUPPORT_ACCELERATE = SetUtils.of(
            "ListBuckets", "CreateBucket", "DeleteBucket");

    // S3Control function getter constant
    private static final String GET_ARN_FROM_INPUT = "getARNMemberValue";
    private static final String UPDATE_ARN_MEMBER_VALUE = "updateARNMemberValue";
    private static final String BACKFILL_ACCOUNT_ID = "backfillAccountID";
    private static final String GET_OUTPOST_ID_FROM_INPUT = "getOutpostIDFromInput";

    // List of operations that use Accesspoint field as ARN input source.
    private  static final Set<String> LIST_ACCESSPOINT_ARN_INPUT = SetUtils.of(
            "GetAccessPoint", "DeleteAccessPoint", "PutAccessPointPolicy",
            "GetAccessPointPolicy", "DeleteAccessPointPolicy"
    );

    // List of operations that use OutpostID to resolve endpoint
    private  static final Set<String> LIST_OUTPOST_ID_INPUT = SetUtils.of(
      "CreateBucket","ListRegionalBuckets"
    );

    /**
     * Gets the sort order of the customization from -128 to 127, with lowest
     * executed first.
     *
     * @return Returns the sort order, defaults to -40.
     */
    @Override
    public byte getOrder() {
        return 127;
    }

    @Override
    public void writeAdditionalFiles(
            GoSettings settings,
            Model model,
            SymbolProvider symbolProvider,
            GoDelegator goDelegator
    ) {
        ServiceShape service = settings.getService(model);

        // if service is s3control
        if (isS3ControlService(model, service)){
            goDelegator.useShapeWriter(service, this::writeS3ControlMiddlewareHelper);
            goDelegator.useShapeWriter(service, writer -> {
                writeARNHelper(writer, model, symbolProvider, service);
                writeOutpostIDHelper(writer, model, symbolProvider, service);
                writeBackfillAccountIDHelper(writer, model, symbolProvider, service);
            });

            for (ShapeId operationID : service.getAllOperations()) {
                OperationShape operation = model.expectShape(operationID, OperationShape.class);
                goDelegator.useShapeWriter(operation, writer -> {
                    writeOpARNHelper(writer, model, symbolProvider, operation);
                    writeOpBackfillAccountIDHelper(writer, model, symbolProvider, operation);
                });
            }
        }

        // check if service is s3
        if (isS3Service(model, service)) {
            goDelegator.useShapeWriter(service, this::writeS3MiddlewareHelper);

            goDelegator.useShapeWriter(service, writer -> {
                writeInputGetter(writer, model, symbolProvider, service);
            });

            goDelegator.useShapeWriter(service, writer -> {
                writeAccelerateValidator(writer, model, symbolProvider, service);
            });
        }
    }


    // writes BackfillAccountID Helper function for s3 api client
    // == api_client.go===
    // func backfillAccountID(in interface, v string) (interface{},error) {
    //      iv, ok := in.(interface{ backfillAccountID(string) (interface{}, error) })
    //      if !ok  { return in, nil } // as account id is not even supported
    //      return iv.backfillAccountID(v)
    // }
    private void writeBackfillAccountIDHelper(GoWriter writer, Model model,SymbolProvider symbolProvider, ServiceShape service) {
     writer.write("func $L(in interface{}, v string) (interface{},error) { ", BACKFILL_ACCOUNT_ID);
     writer.write("iv, ok := in.(interface{ $L(string) (interface{}, error) })", BACKFILL_ACCOUNT_ID);
     writer.write("if !ok  { return in, nil } // as account id is not even supported");
     writer.write("return iv.$L(v) }", BACKFILL_ACCOUNT_ID);
    }

    // writes BackfillAccountID Helper function for s3 api operation
    // === api_operation.go===
    // func (in operationInput) backfillAccountID(v string) (operationInput, error ) {
    //      if in.AccountId!=nil {
    //         iv := *in.AccountId
    //         if !strings.EqualFold(iv, v) {
    //             return &in, fmt.Errorf("error backfilling account id")
    //         }
    //         return &in, nil
    //      }
    //
    //   in.AccountId = &v
    //   return &in, nil
    // }
    private void writeOpBackfillAccountIDHelper(GoWriter writer, Model model,SymbolProvider symbolProvider, OperationShape operation) {
        StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);
        List<MemberShape> targetAccountIDShape = input.getAllMembers().values().stream()
                .filter(m -> m.getMemberName().equals("AccountId"))
                .collect(Collectors.toList());
        // if model has multiple top level shapes targeting `AccountId`, we throw a codegen exception
        if (targetAccountIDShape.size()>1) {
            throw new CodegenException("AccountId shape should be targeted by only one input member, found " +
                    targetAccountIDShape.size() +" for Input shape: "+ input.getId());
        }

        if (!targetAccountIDShape.isEmpty()) {
            Symbol inputSymbol = symbolProvider.toSymbol(input);
            String memberName = targetAccountIDShape.get(0).getMemberName();
            writer.write("func (in $T) $L (v string) (interface{}, error) { ", inputSymbol, BACKFILL_ACCOUNT_ID);
            writer.write("if in.$L != nil {", memberName);

            writer.addUseImports(SmithyGoDependency.STRINGS);
            writer.write("if !strings.EqualFold(*in.$L, v) {", memberName);

            writer.addUseImports(SmithyGoDependency.FMT);
            writer.write("return &in, fmt.Errorf(\"error backfilling account id\") }");
            writer.write("return &in, nil }");
            writer.write("in.$L = &v", memberName);
            writer.write("return &in, nil }");
        }
    }

    // writes getARNMemberValue and updateARNMemberValue update function for all input operations
    private void writeARNHelper(GoWriter writer, Model model,SymbolProvider symbolProvider, ServiceShape service) {
        // write getARNMemberValue
        writer.write("func $L(in interface{}) (*string, bool) { ", GET_ARN_FROM_INPUT);
        writer.write("iv, ok := in.(interface{ $L() (*string, bool) })", GET_ARN_FROM_INPUT);
        writer.write("if !ok  { return nil, false }");
        writer.write("return iv.$L() }", GET_ARN_FROM_INPUT);

        // write UpdateARNMemberValue
        writer.write("func $L(in interface{}, v string) (interface{},bool) { ", UPDATE_ARN_MEMBER_VALUE);
        writer.write("iv, ok := in.(interface{ $L(string) (interface{}) })", UPDATE_ARN_MEMBER_VALUE);
        writer.write("if !ok  { return &in, false }");
        writer.write("return iv.$L(v), true }", UPDATE_ARN_MEMBER_VALUE);
    }

    private void writeOpARNHelper(GoWriter writer, Model model, SymbolProvider symbolProvider, OperationShape operation) {
        // list of outpost id input require special behavior
        if (LIST_OUTPOST_ID_INPUT.contains(operation.getId().getName())) { return; }

        // arn target shape
        String arnType = LIST_ACCESSPOINT_ARN_INPUT.contains(operation.getId().getName())? "AccessPointName":"BucketName";

        StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);
        List<MemberShape> listOfARNMembers = input.getAllMembers().values().stream()
                .filter(m -> m.getTarget().getName().equals(arnType))
                .collect(Collectors.toList());
        // if model has multiple top level shapes targeting arnable field, we throw a codegen exception
        if (listOfARNMembers.size()>1) {
            throw new CodegenException(arnType + " shape should be targeted by only one input member, found " +
                    listOfARNMembers.size() +" for Input shape: "+ input.getId());
        }

        if (!listOfARNMembers.isEmpty()) {
            Symbol inputSymbol = symbolProvider.toSymbol(input);
            String memberName = listOfARNMembers.get(0).getMemberName();
            writer.write("func (in $T) $L () (*string, bool) { ", inputSymbol, GET_ARN_FROM_INPUT);
            writer.write("if in.$L == nil {return nil, false }", memberName);
            writer.write("return in.$L, true }", memberName);

            writer.insertTrailingNewline();
            writer.write("func (in $T) $L (v string) interface{} {", inputSymbol, UPDATE_ARN_MEMBER_VALUE);
            writer.write("in.$L = &v", memberName);
            writer.write("return &in }");
        }
    }


    // writes OutpostID Helper function for operations CreateBucket and ListRegionalBuckets
    //
    //  func getOutpostIDFromInput (in interface{}) (*string, bool) {
    //      switch in.(type) {
    //          case CreateBucket: return in.OutpostId, true
    //          case listRegionalBuckets : return in.OutpostID, true
    //          default: nil, false
    //      }
    // }
    private void writeOutpostIDHelper(GoWriter writer, Model model,SymbolProvider symbolProvider, ServiceShape service) {
        writer.writeDocs("getOutpostIDFromInput returns a boolean indicating if the input has a modeled outpost-id, " +
                " and a pointer to string denoting a provided outpost-id member value");
        writer.openBlock("func $L (input interface{}) (*string, bool) {","}",
                GET_OUTPOST_ID_FROM_INPUT, ()-> {
            writer.openBlock("switch i:= input.(type) {", "}", () -> {
                for (ShapeId operationId : service.getAllOperations()) {
                    // customization only applied to operations CreateBucket, ListRegionalBuckets
                    if (!LIST_OUTPOST_ID_INPUT.contains(operationId.getName())) {
                        continue;
                    }
                    OperationShape operation = model.expectShape(operationId, OperationShape.class);
                    StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);
                    List<MemberShape> outpostIDMemberShapes = input.getAllMembers().values().stream()
                            .filter(m -> m.getMemberName().equalsIgnoreCase("OutpostId"))
                            .collect(Collectors.toList());
                    // if model has multiple top level shapes targeting `OutpostId`, we throw a codegen exception
                    if (outpostIDMemberShapes.size() > 1) {
                        throw new CodegenException(
                                "OutpostID shape should be targeted by only one input member, found " +
                                        outpostIDMemberShapes.size() + " for Input shape: " + input.getId());
                    }

                    if (!outpostIDMemberShapes.isEmpty()) {
                        writer.write("case $P: return i.$L, true", symbolProvider.toSymbol(input),
                                outpostIDMemberShapes.get(0).getMemberName());
                    }
                }
                writer.write("default: return nil, false");
            });
        });
    }


    private void writeAccelerateValidator(GoWriter writer, Model model, SymbolProvider symbolProvider, ServiceShape service) {
        writer.writeDocs("supportAccelerate returns a boolean indicating if the operation associated with the provided input "
                + "supports S3 Transfer Acceleration");
        writer.openBlock("func $L(input interface{}) bool {", "}", SUPPORT_ACCELERATE, () -> {
            writer.openBlock("switch input.(type) {" , "}", () -> {
                for (ShapeId operationId : service.getAllOperations()) {
                    // check if operation does not support s3 accelerate
                    if (NOT_SUPPORT_ACCELERATE.contains(operationId.getName())) {
                        OperationShape operation = model.expectShape(operationId, OperationShape.class);
                        StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);
                        writer.write("case $P: return false", symbolProvider.toSymbol(input));
                    }
                }
               writer.write("default: return true");
            });
        });
    }


    private void writeInputGetter(GoWriter writer, Model model, SymbolProvider symbolProvider, ServiceShape service) {
        writer.writeDocs("getBucketFromInput returns a boolean indicating if the input has a modeled bucket name, " +
                " and a pointer to string denoting a provided bucket member value");
        writer.openBlock("func getBucketFromInput(input interface{}) (*string, bool) {","}", ()-> {
            writer.openBlock("switch i:= input.(type) {", "}", () -> {
                service.getAllOperations().forEach((operationId)-> {
                    OperationShape operation = model.expectShape(operationId, OperationShape.class);
                    StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);

                    List<MemberShape> targetBucketShape = input.getAllMembers().values().stream()
                            .filter(m -> m.getTarget().getName().equals("BucketName"))
                            .collect(Collectors.toList());
                    // if model has multiple top level shapes targeting `BucketName`, we throw a codegen exception
                    if (targetBucketShape.size()>1) {
                        throw new CodegenException("BucketName shape should be targeted by only one input member, found " +
                                targetBucketShape.size() +" for Input shape: "+ input.getId());
                    }

                    if (!targetBucketShape.isEmpty() && !operationId.getName().equalsIgnoreCase("GetBucketLocation")) {
                        writer.write("case $P: return i.$L, true", symbolProvider.toSymbol(input), targetBucketShape.get(0).getMemberName());
                    }
                });
                writer.write("default: return nil, false");
            });
        });
    }

    private void writeS3ControlMiddlewareHelper(GoWriter writer) {
        // imports
        writer.addUseImports(SmithyGoDependency.SMITHY_MIDDLEWARE);
        writer.addUseImports(AwsCustomGoDependency.S3CONTROL_CUSTOMIZATION);

        writer.openBlock("func $L(stack *middleware.Stack, options Options) {", "}", UPDATE_ENDPOINT_ADDER, () -> {
            writer.write("$T(stack, $T{ \n"
                            + "GetARNInput: $L,\n BackfillAccountID: $L,\n GetOutpostIDInput: $L, \n "
                            + "UpdateARNField: $L,\n EndpointResolver: options.EndpointResolver,\n "
                            + "EndpointResolverOptions: options.EndpointOptions,\n"
                            + "UseDualstack: options.$L, \n UseARNRegion: options.$L, \n })",
                    SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_ADDER,
                            AwsCustomGoDependency.S3CONTROL_CUSTOMIZATION).build(),
                    SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_ADDER + "Options",
                            AwsCustomGoDependency.S3CONTROL_CUSTOMIZATION).build(),
                    GET_ARN_FROM_INPUT,
                    BACKFILL_ACCOUNT_ID,
                    GET_OUTPOST_ID_FROM_INPUT,
                    UPDATE_ARN_MEMBER_VALUE,
                    USE_DUALSTACK_OPTION,
                    USE_ARNREGION_OPTION
            );
        });
        writer.insertTrailingNewline();
    }

    private void writeS3MiddlewareHelper(GoWriter writer) {
        // imports
        writer.addUseImports(SmithyGoDependency.SMITHY_MIDDLEWARE);
        writer.addUseImports(AwsCustomGoDependency.S3_CUSTOMIZATION);

        writer.openBlock("func $L(stack *middleware.Stack, options Options) {", "}", UPDATE_ENDPOINT_ADDER, () -> {
            writer.write("$T(stack, $T{ \n"
                            + "GetBucketFromInput: $L,\n UsePathStyle: options.$L,\n "
                            + "UseAccelerate: options.$L,\n SupportsAccelerate: $L,\n "
                            + "EndpointResolver: options.EndpointResolver,\n "
                            + "EndpointResolverOptions: options.EndpointOptions,\n"
                            + "UseDualstack: options.$L, \n UseARNRegion: options.$L, \n })",
                    SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_ADDER,
                            AwsCustomGoDependency.S3_CUSTOMIZATION).build(),
                    SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_ADDER + "Options",
                            AwsCustomGoDependency.S3_CUSTOMIZATION).build(),
                    GET_BUCKET_FROM_INPUT,
                    USE_PATH_STYLE_OPTION,
                    USE_ACCELERATE_OPTION,
                    SUPPORT_ACCELERATE,
                    USE_DUALSTACK_OPTION,
                    USE_ARNREGION_OPTION
            );
        });
        writer.insertTrailingNewline();
    }

    @Override
    public List<RuntimeClientPlugin> getClientPlugins() {
        return ListUtils.of(
                // Add S3 config to use path style host addressing.
                RuntimeClientPlugin.builder()
                        .servicePredicate(S3UpdateEndpoint::isS3Service)
                        .configFields(ListUtils.of(
                                ConfigField.builder()
                                        .name(USE_PATH_STYLE_OPTION)
                                        .type(SymbolUtils.createValueSymbolBuilder("bool")
                                                .putProperty(SymbolUtils.GO_UNIVERSE_TYPE, true)
                                                .build())
                                        .documentation("Allows you to enable the client to use path-style addressing, "
                                                + "i.e., `https://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client "
                                                + "will use virtual hosted bucket addressing when possible"
                                                + "(`https://BUCKET.s3.amazonaws.com/KEY`).")
                                        .build(),
                                ConfigField.builder()
                                        .name(USE_ACCELERATE_OPTION)
                                        .type(SymbolUtils.createValueSymbolBuilder("bool")
                                                .putProperty(SymbolUtils.GO_UNIVERSE_TYPE, true)
                                                .build())
                                        .documentation("Allows you to enable S3 Accelerate feature. All operations "
                                                + "compatible with S3 Accelerate will use the accelerate endpoint for "
                                                + "requests. Requests not compatible will fall back to normal S3 requests. "
                                                + "The bucket must be enabled for accelerate to be used with S3 client with "
                                                + "accelerate enabled. If the bucket is not enabled for accelerate an error "
                                                + "will be returned. The bucket name must be DNS compatible to work "
                                                + "with accelerate.")
                                        .build()
                                ))
                        .build(),
                // Add S3 shared config's dualstack option
                RuntimeClientPlugin.builder()
                        .servicePredicate(S3UpdateEndpoint::isS3SharedService)
                        .configFields(ListUtils.of(
                            ConfigField.builder()
                            .name(USE_DUALSTACK_OPTION)
                            .type(SymbolUtils.createValueSymbolBuilder("bool")
                                .putProperty(SymbolUtils.GO_UNIVERSE_TYPE, true)
                                .build())
                            .documentation("Allows you to enable Dualstack endpoint support for the service.")
                            .build(),
                                ConfigField.builder()
                                        .name(USE_ARNREGION_OPTION)
                                        .type(SymbolUtils.createValueSymbolBuilder("bool")
                                                .putProperty(SymbolUtils.GO_UNIVERSE_TYPE, true)
                                                .build())
                                        .documentation("Allows you to enable arn region support for the service.")
                                        .build()
                        ))
                        .registerMiddleware(MiddlewareRegistrar.builder()
                                .resolvedFunction(SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_ADDER).build())
                                .useClientOptions()
                                .build())
                .build()
        );
    }

    private static boolean isS3Service(Model model, ServiceShape service) {
        return service.expectTrait(ServiceTrait.class).getSdkId().equalsIgnoreCase("S3");
    }

    private static boolean isS3ControlService(Model model, ServiceShape service) {
        return service.expectTrait(ServiceTrait.class).getSdkId().equalsIgnoreCase("S3 Control");
    }

    private static boolean isS3SharedService(Model model, ServiceShape service) {
        return isS3Service(model,service) || isS3ControlService(model, service);
    }
}
