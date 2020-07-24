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
 */

package software.amazon.smithy.aws.go.codegen;

import static software.amazon.smithy.go.codegen.integration.ProtocolUtils.writeSafeMemberAccessor;

import java.util.Optional;
import java.util.Set;
import java.util.TreeSet;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.SyntheticClone;
import software.amazon.smithy.go.codegen.integration.HttpProtocolTestGenerator;
import software.amazon.smithy.go.codegen.integration.HttpProtocolUnitTestGenerator;
import software.amazon.smithy.go.codegen.integration.HttpProtocolUnitTestRequestGenerator;
import software.amazon.smithy.go.codegen.integration.HttpProtocolUnitTestResponseErrorGenerator;
import software.amazon.smithy.go.codegen.integration.HttpProtocolUnitTestResponseGenerator;
import software.amazon.smithy.go.codegen.integration.IdempotencyTokenMiddlewareGenerator;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator.GenerationContext;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.traits.XmlAttributeTrait;
import software.amazon.smithy.model.traits.XmlNameTrait;
import software.amazon.smithy.model.traits.XmlNamespaceTrait;
import software.amazon.smithy.utils.SetUtils;

/**
 * Utility methods for generating AWS protocols.
 */
final class AwsProtocolUtils {
    private AwsProtocolUtils() {
    }

    /**
     * Generates HTTP protocol tests with all required AWS-specific configuration set.
     *
     * @param context The generation context.
     */
    static void generateHttpProtocolTests(GenerationContext context) {
        Set<HttpProtocolUnitTestGenerator.ConfigValue> configValues = new TreeSet<>(SetUtils.of(
                HttpProtocolUnitTestGenerator.ConfigValue.builder()
                        .name(AddAwsConfigFields.REGION_CONFIG_NAME)
                        .value(writer -> writer.write("$S,", "us-west-2"))
                        .build(),
                HttpProtocolUnitTestGenerator.ConfigValue.builder()
                        .name(AddAwsConfigFields.HTTP_CLIENT_CONFIG_NAME)
                        .value(writer -> {
                            writer.addUseImports(AwsGoDependency.AWS_CORE);
                            writer.write("aws.NewBuildableHTTPClient(),");
                        })
                        .build(),
                HttpProtocolUnitTestGenerator.ConfigValue.builder()
                        .name(AddAwsConfigFields.ENDPOINT_RESOLVER_CONFIG_NAME)
                        .value(writer -> {
                            writer.addUseImports(AwsGoDependency.AWS_CORE);
                            writer.openBlock("aws.EndpointResolverFunc("
                                            + "func(service, region string) (e aws.Endpoint, err error) {",
                                    "}),", () -> {
                                        writer.write("e.URL = server.URL");
                                        writer.write("e.SigningRegion = \"us-west-2\"");
                                        writer.write("return e, err");
                                    });
                        })
                        .build(),
                HttpProtocolUnitTestGenerator.ConfigValue.builder()
                        .name("APIOptions")
                        .value(writer -> {
                            writer.addUseImports(SmithyGoDependency.SMITHY_MIDDLEWARE);
                            writer.openBlock("[]APIOptionFunc{", "},", () -> {
                                writer.openBlock("func(s *middleware.Stack) error {", "},", () -> {
                                    writer.write("s.Finalize.Clear()");
                                    writer.write("return nil");
                                });
                            });
                        })
                        .build()
        ));

        // TODO can this check be replaced with a lookup into the runtime plugins?
        if (IdempotencyTokenMiddlewareGenerator.hasOperationsWithIdempotencyToken(context.getModel(),
                context.getService())) {
            configValues.add(
                    HttpProtocolUnitTestGenerator.ConfigValue.builder()
                            .name(IdempotencyTokenMiddlewareGenerator.IDEMPOTENCY_CONFIG_NAME)
                            .value(writer -> {
                                writer.addUseImports(SmithyGoDependency.SMITHY_RAND);
                                writer.addUseImports(SmithyGoDependency.SMITHY_TESTING);
                                writer.write("smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),");
                            })
                            .build()
            );
        }

        new HttpProtocolTestGenerator(context,
                (HttpProtocolUnitTestRequestGenerator.Builder) new HttpProtocolUnitTestRequestGenerator
                        .Builder()
                        .addClientConfigValues(configValues),
                (HttpProtocolUnitTestResponseGenerator.Builder) new HttpProtocolUnitTestResponseGenerator
                        .Builder()
                        .addClientConfigValues(configValues),
                (HttpProtocolUnitTestResponseErrorGenerator.Builder) new HttpProtocolUnitTestResponseErrorGenerator
                        .Builder()
                        .addClientConfigValues(configValues)
        ).generateProtocolTests();
    }

    public static void writeJsonErrorMessageCodeDeserializer(GenerationContext context) {
        GoWriter writer = context.getWriter();
        // The error code could be in the headers, even though for this protocol it should be in the body.
        writer.write("code := response.Header.Get(\"X-Amzn-ErrorType\")");
        writer.write("if len(code) != 0 { errorCode = restjson.SanitizeErrorCode(code) }");
        writer.write("");

        initializeJsonDecoder(writer, "errorBody");
        writer.addUseImports(AwsGoDependency.AWS_REST_JSON_PROTOCOL);
        // This will check various body locations for the error code and error message
        writer.write("code, message, err := restjson.GetErrorInfo(decoder)");
        handleDecodeError(writer);

        writer.addUseImports(SmithyGoDependency.IO);
        // Reset the body in case it needs to be used for anything else.
        writer.write("errorBody.Seek(0, io.SeekStart)");

        // Only set the values if something was found so that we keep the default values.
        writer.write("if len(code) != 0 { errorCode = restjson.SanitizeErrorCode(code) }");
        writer.write("if len(message) != 0 { errorMessage = message }");
        writer.write("");
    }

    public static void initializeJsonDecoder(GoWriter writer, String bodyLocation) {
        // Use a ring buffer and tee reader to help in pinpointing any deserialization errors.
        writer.addUseImports(SmithyGoDependency.SMITHY_IO);
        writer.write("buff := make([]byte, 1024)");
        writer.write("ringBuffer := smithyio.NewRingBuffer(buff)");
        writer.write("");

        writer.addUseImports(SmithyGoDependency.IO);
        writer.addUseImports(SmithyGoDependency.JSON);
        writer.write("body := io.TeeReader($L, ringBuffer)", bodyLocation);
        writer.write("decoder := json.NewDecoder(body)");
        writer.write("decoder.UseNumber()");
        writer.write("");
    }

    public static void handleDecodeError(GoWriter writer, String returnExtras) {
        writer.openBlock("if err != nil {", "}", () -> {
            writer.addUseImports(SmithyGoDependency.BYTES);
            writer.addUseImports(SmithyGoDependency.SMITHY);
            writer.write("var snapshot bytes.Buffer");
            writer.write("io.Copy(&snapshot, ringBuffer)");
            writer.openBlock("return $L&smithy.DeserializationError {", "}", returnExtras, () -> {
                writer.write("Err: fmt.Errorf(\"failed to decode response body with invalid JSON, %w\", err),");
                writer.write("Snapshot: snapshot.Bytes(),");
            });
        }).write("");
    }

    public static void handleDecodeError(GoWriter writer) {
        handleDecodeError(writer, "");
    }

    /**
     * generateXMLStartElement generates the XML start element for a shape. It is used to generate smithy xml's startElement.
     *
     * @param context is the generation context.
     * @param shape is the Shape for which xml start element is to be generated.
     * @param dst is the operand name which holds the generated start element.
     * @param inputSrc is the input variable for the shape with values to be serialized.
     */
    public static void generateXMLStartElement(GenerationContext context, Shape shape, String dst, String inputSrc) {
        GoWriter writer = context.getWriter();
        String attrName = dst + "Attr";
        writer.write("$L := []smithyxml.Attr{}", attrName);

        Optional<XmlNamespaceTrait> xmlNamespaceTrait = shape.getTrait(XmlNamespaceTrait.class);
        if (xmlNamespaceTrait.isPresent()) {
            XmlNamespaceTrait ns = xmlNamespaceTrait.get();
            writer.write("$L = append($L, smithyxml.NewNamespaceAttribute($S, $S))",
                    attrName, attrName,
                    ns.getPrefix().isPresent() ? ns.getPrefix().get() : "", ns.getUri()
            );
        }

        // Traverse member shapes to get attributes
        shape.members().stream().forEach(memberShape -> {
            Optional<XmlAttributeTrait> xmlAttributeTrait = memberShape.getTrait(XmlAttributeTrait.class);
            if (xmlAttributeTrait.isPresent()) {
                writeSafeMemberAccessor(context, memberShape, inputSrc, (operand) -> {
                    writer.write("$L = append($L, smithyxml.NewAttribute($S, string(*$L)))",
                            attrName, attrName,
                            getSerializedXMLMemberName(memberShape), operand);
                });
            }
        });

        writer.openBlock("$L := smithyxml.StartElement{ ", "}", dst, () -> {
            writer.openBlock("Name:smithyxml.Name{", "},", () -> {
                writer.write("Local: $S,", getSerializedXMLShapeName(context, shape));
            });
            writer.write("Attr : $L,", attrName);
        });
    }


    /**
     * getSerializedXMLMemberName returns a xml member name used for serializing. If a member shape has
     * XML name trait, xml name would be given precedence over member name.
     *
     * @param memberShape is the member shape for which serializer name is queried.
     * @return name of a xml member shape used by serializers
     */
    private static String getSerializedXMLMemberName(MemberShape memberShape) {
        Optional<XmlNameTrait> xmlNameTrait = memberShape.getTrait(XmlNameTrait.class);
        return xmlNameTrait.isPresent() ? xmlNameTrait.get().getValue() : memberShape.getMemberName();
    }

    /**
     * getSerializedXMLShapeName returns a xml shape name used for serializing. If a member shape
     * has xml name trait, xml name would be given precedence over member name.
     * This correctly handles renamed shapes, and returns the original shape name.
     *
     * @param context is the generation context for which
     * @param shape is the Shape for which serializer name is queried.
     * @return name of a xml member shape used by serializers.
     */
    private static String getSerializedXMLShapeName(GenerationContext context, Shape shape) {
        SymbolProvider symbolProvider = context.getSymbolProvider();
        Symbol shapeSymbol = symbolProvider.toSymbol(shape);
        String shapeName = shapeSymbol.getName();

        // check if synthetic cloned shape
        Optional<SyntheticClone> clone = shape.getTrait(SyntheticClone.class);
        if (clone.isPresent()) {
            SyntheticClone cl = clone.get();
            shapeName = cl.getArchetype().getName();
        }

        // check if shape is member shape
        Optional<MemberShape> member = shape.asMemberShape();
        if (member.isPresent()) {
            MemberShape memberShape = member.get();
            return getSerializedXMLMemberName(memberShape);
        }

        Optional<XmlNameTrait> xmlNameTrait = shape.getTrait(XmlNameTrait.class);
        return xmlNameTrait.isPresent() ? xmlNameTrait.get().getValue() : shapeName;
    }
}
