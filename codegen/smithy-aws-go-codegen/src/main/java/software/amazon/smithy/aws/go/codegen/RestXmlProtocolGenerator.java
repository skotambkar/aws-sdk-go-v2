package software.amazon.smithy.aws.go.codegen;

import static software.amazon.smithy.aws.go.codegen.AwsProtocolUtils.handleDecodeError;
import static software.amazon.smithy.aws.go.codegen.AwsProtocolUtils.initializeJsonDecoder;
import static software.amazon.smithy.aws.go.codegen.AwsProtocolUtils.initializeXMLDecoder;

import java.util.Set;
import java.util.TreeSet;
import java.util.function.Predicate;
import java.util.stream.Collectors;
import software.amazon.smithy.aws.traits.protocols.RestXmlTrait;
import software.amazon.smithy.codegen.core.CodegenException;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.go.codegen.ApplicationProtocol;
import software.amazon.smithy.go.codegen.GoStackStepMiddlewareGenerator;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.integration.HttpBindingProtocolGenerator;
import software.amazon.smithy.go.codegen.integration.HttpProtocolGeneratorUtils;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator;
import software.amazon.smithy.go.codegen.integration.ProtocolUtils;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.knowledge.HttpBinding;
import software.amazon.smithy.model.knowledge.HttpBindingIndex;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.OperationShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.model.traits.EnumTrait;
import software.amazon.smithy.model.traits.StreamingTrait;
import software.amazon.smithy.model.traits.TimestampFormatTrait;

abstract class RestXmlProtocolGenerator extends HttpBindingProtocolGenerator {
    /**
     * Creates a AWS REST XML protocol generator.
     */
    RestXmlProtocolGenerator() {
        super(true);
    }

    @Override
    public void generateProtocolTests(GenerationContext context) {
        AwsProtocolUtils.generateHttpProtocolTests(context);
    }

    @Override
    protected TimestampFormatTrait.Format getDocumentTimestampFormat() {
        return null;
    }

    @Override
    protected String getDocumentContentType() {
        return null;
    }

    @Override
    protected void generateOperationDocumentSerializer(GenerationContext context, OperationShape operation) {

    }

    @Override
    protected void writeMiddlewareDocumentSerializerDelegator(
            GenerationContext context,
            OperationShape operation,
            GoStackStepMiddlewareGenerator generator
    ) {

    }

    @Override
    protected void writeMiddlewarePayloadSerializerDelegator(
            GenerationContext context,
            OperationShape operation,
            MemberShape memberShape,
            GoStackStepMiddlewareGenerator generator
    ) {

    }

    @Override
    protected void generateDocumentBodyShapeSerializers(
            GenerationContext context, Set<Shape> shapes
    ) {

    }

    /*     ================Deserializer===========================     */

    /*======================== Error Deser ================================*/


    @Override
    public void handleOperationErrorResponse(GenerationContext context, OperationShape operation) {
        GoWriter writer = context.getWriter();
        String errorFunctionName = ProtocolGenerator.getOperationErrorDeserFunctionName(
                operation, context.getProtocolName());

        writer.addUseImports(SmithyGoDependency.SMITHY_IO);
        writer.write("buff := make([]byte, 1024)");
        writer.write("ringBuffer := smithyio.NewRingBuffer(buff)");
//        writer.write("var errorBody *bytes.Buffer");
//        writer.write("w := io.MultiWriter(ringBuffer, errorBody)");
        writer.write("");

        writer.addUseImports(SmithyGoDependency.IO);
        writer.write("body := io.TeeReader(response.Body, ringBuffer)");
        writer.write("_  = body");
        writer.write("defer response.Body.Close()");
        writer.write("");

        writer.addUseImports(SmithyGoDependency.SMITHY_DECODING);
        writer.write("var decoder *smithydecoding.XMLNodeDecoder");
        writer.write("_ = decoder");

        // TODO: Add customization for S3 to check for response status code 200
        writer.openBlock("if response.StatusCode < 200 || response.StatusCode >= 300 {", "}", () -> {
//            writer.addUseImports(SmithyGoDependency.XML);
//            writer.write("rootDecoder := xml.NewDecoder(body)");
//
//            writer.write("// fetch the root element ignoring comments and preamble");
//            writer.write("t, err  := smithydecoding.RootElement(rootDecoder)");
//
//            writer.addUseImports(SmithyGoDependency.IO);
//            writer.write("if err == io.EOF { err = nil }");
//            writer.write("if err != nil {return out, metadata, "
//                    + "fmt.Errorf(\"error fetching the start element of xml response body: %w\", err)}");
//            writer.write("decoder = smithydecoding.NewXMLNodeDecoder(rootDecoder, t)");
//            writer.insertTrailingNewline();

            writer.write("return out, metadata, $L(response)", errorFunctionName);
        });
    }

    @Override
    public void generateErrorDeserializer(GenerationContext context, StructureShape shape) {
        GoWriter writer = context.getWriter();
        String functionName = ProtocolGenerator.getErrorDeserFunctionName(shape, context.getProtocolName());
        Symbol responseType = getApplicationProtocol().getResponseType();

        writer.addUseImports(SmithyGoDependency.SMITHY_DECODING);

        writer.openBlock("func $L(decoder *smithydecoding.XMLNodeDecoder, response $P) error {", "}",
                functionName, responseType, () -> deserializeError(context, shape));
        writer.insertTrailingNewline();
    }

    protected void deserializeError(GenerationContext context, StructureShape shape) {
        GoWriter writer = context.getWriter();
        Symbol symbol = context.getSymbolProvider().toSymbol(shape);

        writer.write("output := &$T{}", symbol);
        writer.write("_ = output");
        writer.write("");

        // TODO: filter on error document body contains
        if (isShapeWithResponseBindings(context.getModel(), shape, HttpBinding.Location.DOCUMENT)) {
            String documentDeserFunctionName = ProtocolGenerator.getDocumentDeserializerFunctionName(
                    shape, getProtocolName());

            writer.write("err := $L(&output, decoder)", documentDeserFunctionName);
            // TODO follow json pattern and return wrapped deserialization Error
            writer.addUseImports(SmithyGoDependency.IO);
            writer.write("if err == io.EOF { err = nil }");
            writer.write("if err != nil {return err}");
            writer.write("");
        }

        if (isShapeWithRestResponseBindings(context.getModel(), shape)) {
            String bindingDeserFunctionName = ProtocolGenerator.getOperationHttpBindingsDeserFunctionName(
                    shape, getProtocolName());
            writer.openBlock("if err := $L(output, response); err != nil {", "}", bindingDeserFunctionName, () -> {
                writer.addUseImports(SmithyGoDependency.SMITHY);
                writer.write(String.format("return &smithy.DeserializationError{Err: %s}",
                        "fmt.Errorf(\"failed to decode response error with invalid HTTP bindings, %w\", err)"));
            });
            writer.write("");
        }

        writer.write("return output");
    }

    @Override
    public void writeErrorMessageCodeDeserializer(GenerationContext context, OperationShape operation) {
        GoWriter writer = context.getWriter();
        boolean isNoErrorWrapping = false;
        if (operation.hasTrait(RestXmlTrait.ID)) {
            RestXmlTrait trait = operation.getTrait(RestXmlTrait.class).get();
            isNoErrorWrapping = trait.isNoErrorWrapping();
        }
        writer.write("errorCode, err := smithydecoding.GetXMLResponseErrorCode(errorBody, $L)", isNoErrorWrapping);
        writer.write("if err != nil { return err}");
        writer.insertTrailingNewline();

        writer.write("errorBody.Seek(0, io.SeekStart)");
        writer.insertTrailingNewline();
    }

    @Override
    public Set<StructureShape> generateErrorDispatcher(GenerationContext context, OperationShape operation){
        ApplicationProtocol applicationProtocol = getApplicationProtocol();
        Symbol responseType = applicationProtocol.getResponseType();
        return HttpProtocolGeneratorUtils.generateXmlErrorDispatcher(
                context, operation, responseType, (c, s) -> writeErrorMessageCodeDeserializer(c, s));
    }

    /*======================== Document Deser ================================*/
    @Override
    protected void writeMiddlewareDocumentDeserializerDelegator(
            GenerationContext context,
            OperationShape operation,
            GoStackStepMiddlewareGenerator generator
    ) {
        Model model = context.getModel();
        GoWriter writer = context.getWriter();
        Shape targetShape = ProtocolUtils.expectOutput(model, operation);
        String operand = "output";

        boolean isShapeWithPayloadBinding = isShapeWithResponseBindings(model, operation, HttpBinding.Location.PAYLOAD);
        if (isShapeWithPayloadBinding) {
            // since payload trait can only be applied to a single member in a output shape
            MemberShape memberShape = model.getKnowledge(HttpBindingIndex.class)
                    .getResponseBindings(operation, HttpBinding.Location.PAYLOAD).stream()
                    .findFirst()
                    .orElseThrow(() -> new CodegenException("Expected payload binding member"))
                    .getMember();

            Shape payloadShape = model.expectShape(memberShape.getTarget());

            // if target shape is of type String or type Blob, then delegate deserializers for explicit payload shapes
            if (payloadShape.isStringShape() || payloadShape.isBlobShape()) {
                writeMiddlewarePayloadBindingDeserializerDelegator(writer, targetShape);
                return;
            }
            // for other payload target types we should deserialize using the appropriate document deserializer
            targetShape = payloadShape;
            operand += "." + context.getSymbolProvider().toMemberName(memberShape);
        }

        writeMiddlewareDocumentBindingDeserializerDelegator(writer, targetShape, operand);
    }

    @Override
    protected void generateOperationDocumentDeserializer(
            GenerationContext context, OperationShape operation
    ) {
        Model model = context.getModel();
        HttpBindingIndex bindingIndex = model.getKnowledge(HttpBindingIndex.class);
        Set<MemberShape> documentBindings = bindingIndex.getResponseBindings(operation, HttpBinding.Location.DOCUMENT)
                .stream()
                .map(HttpBinding::getMember)
                .collect(Collectors.toSet());

        Shape outputShape = ProtocolUtils.expectOutput(model, operation);
        GoWriter writer = context.getWriter();

        if (documentBindings.size() != 0) {
            outputShape.accept(new XmlShapeDeserVisitor(context, documentBindings::contains));
        }

        Set<MemberShape> payloadBindings = bindingIndex.getResponseBindings(operation, HttpBinding.Location.PAYLOAD)
                .stream()
                .map(HttpBinding::getMember)
                .collect(Collectors.toSet());

        if (payloadBindings.size() == 0) {
            return;
        }

        writePayloadBindingDeserializer(context, outputShape, payloadBindings::contains);
        writer.write("");
    }

    @Override
    protected void generateDocumentBodyShapeDeserializers(
            GenerationContext context, Set<Shape> shapes
    ) {
        XmlShapeDeserVisitor visitor = new XmlShapeDeserVisitor(context);
//        shapes.forEach(shape -> shape.accept(visitor));
        for (Shape shape : shapes) {
            if (shape.isMapShape()) {
                visitor.generateFlattenedMap(context, shape.asMapShape().get());
            }

            if (shape.isListShape()) {
                visitor.generateFlattenedCollection(context, shape.asListShape().get());
            }

            if (shape.isSetShape()) {
                visitor.generateFlattenedCollection(context, shape.asSetShape().get());
            }

            shape.accept(visitor);
        }
    }

    // Generate deserializers for shapes with payload binding
    private void writePayloadBindingDeserializer(
            GenerationContext context,
            Shape shape,
            Predicate<MemberShape> filterMemberShapes
    ) {
        GoWriter writer = context.getWriter();
        SymbolProvider symbolProvider = context.getSymbolProvider();
        Symbol shapeSymbol = symbolProvider.toSymbol(shape);
        String funcName = ProtocolGenerator.getDocumentDeserializerFunctionName(shape, getProtocolName());

        for (MemberShape memberShape : new TreeSet<>(shape.members())) {
            if (!filterMemberShapes.test(memberShape)) {
                continue;
            }

            String memberName = symbolProvider.toMemberName(memberShape);
            Shape targetShape = context.getModel().expectShape(memberShape.getTarget());
            if (targetShape.isStringShape() || targetShape.isBlobShape()) {
                writer.openBlock("func $L(v $P, body io.ReadCloser) error {", "}",
                        funcName, shapeSymbol, () -> {
                            writer.openBlock("if v == nil {", "}", () -> {
                                writer.write("return fmt.Errorf(\"unsupported deserialization of nil %T\", v)");
                            });
                            writer.write("");

                            if (!targetShape.hasTrait(StreamingTrait.class)) {
                                writer.addUseImports(SmithyGoDependency.IOUTIL);
                                writer.write("bs, err := ioutil.ReadAll(body)");
                                writer.write("if err != nil { return err }");
                                writer.openBlock("if len(bs) > 0 {", "}", () -> {
                                    if (targetShape.isBlobShape()) {
                                        writer.write("v.$L = bs", memberName);
                                    } else { // string
                                        writer.addUseImports(SmithyGoDependency.SMITHY_PTR);
                                        if (targetShape.hasTrait(EnumTrait.class)) {
                                            writer.write("v.$L = string(bs)", memberName);
                                        } else {
                                            writer.write("v.$L = ptr.String(string(bs))", memberName);
                                        }
                                    }
                                });
                            } else {
                                writer.write("v.$L = body", memberName);
                            }
                            writer.write("return nil");
                        });
            } else {
                shape.accept(new XmlShapeDeserVisitor(context, filterMemberShapes));
            }
        }
    }

    // Writes middleware that delegates to deserializers for shapes that have explicit payload.
    private void writeMiddlewarePayloadBindingDeserializerDelegator(GoWriter writer, Shape shape) {
        String deserFuncName = ProtocolGenerator.getDocumentDeserializerFunctionName(shape, getProtocolName());
        writer.write("err = $L(output, response.Body)", deserFuncName);
        writer.openBlock("if err != nil {", "}", () -> {
            writer.addUseImports(SmithyGoDependency.SMITHY);
            writer.write(String.format("return out, metadata, &smithy.DeserializationError{Err:%s}",
                    "fmt.Errorf(\"failed to deserialize response payload, %w\", err)"));
        });
    }

    // Write middleware that delegates to deserializers for shapes that have implicit payload
    private void writeMiddlewareDocumentBindingDeserializerDelegator(
            GoWriter writer,
            Shape shape,
            String operand
    ) {
        writer.openBlock("if decoder == nil {", "}", () -> {
            writer.addUseImports(SmithyGoDependency.XML);
            writer.addUseImports(SmithyGoDependency.SMITHY_DECODING);
            writer.write("rootDecoder := xml.NewDecoder(body)");

            writer.write("// fetch the root element ignoring comments and preamble");
            writer.write("t, err  := smithydecoding.RootElement(rootDecoder)");

            writer.addUseImports(SmithyGoDependency.IO);
            writer.write("if err == io.EOF { err = nil }");
            writer.write("if err != nil {return out, metadata, "
                    + "fmt.Errorf(\"error fetching the start element of xml response body: %w\", err)}");
            writer.write("decoder = smithydecoding.NewXMLNodeDecoder(rootDecoder, t)");
            writer.insertTrailingNewline();
        });

        String deserFuncName = ProtocolGenerator.getDocumentDeserializerFunctionName(shape, getProtocolName());
        writer.write("err = $L(&$L, decoder)", deserFuncName, operand);

        writer.addUseImports(SmithyGoDependency.IO);
        writer.openBlock("if err != nil {", "}", () -> {
            writer.addUseImports(SmithyGoDependency.BYTES);
            writer.addUseImports(SmithyGoDependency.SMITHY);
            writer.write("var snapshot bytes.Buffer");
            writer.write("io.Copy(&snapshot, ringBuffer)");
            writer.openBlock("return out, metadata, &smithy.DeserializationError {", "}", () -> {
                writer.write("Err: fmt.Errorf(\"failed to decode response body with invalid XML, %w\", err),");
                writer.write("Snapshot: snapshot.Bytes(),");
            });
        });
    }
}
