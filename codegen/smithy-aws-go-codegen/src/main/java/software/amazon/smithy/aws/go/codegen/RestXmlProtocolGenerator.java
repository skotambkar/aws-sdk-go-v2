package software.amazon.smithy.aws.go.codegen;

import static software.amazon.smithy.go.codegen.integration.ProtocolUtils.writeSafeMemberAccessor;

import java.util.Optional;
import java.util.Set;
import java.util.stream.Collectors;
import software.amazon.smithy.go.codegen.GoStackStepMiddlewareGenerator;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.integration.HttpBindingProtocolGenerator;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator;
import software.amazon.smithy.go.codegen.integration.ProtocolUtils;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.knowledge.HttpBinding;
import software.amazon.smithy.model.knowledge.HttpBindingIndex;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.OperationShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.shapes.ShapeId;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.model.traits.MediaTypeTrait;
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
    public void generateSharedDeserializerComponents(GenerationContext context) {
        super.generateSharedDeserializerComponents(context);
    }

    @Override
    public void generateProtocolTests(GenerationContext context) {
        AwsProtocolUtils.generateHttpProtocolTests(context);
    }

    @Override
    protected TimestampFormatTrait.Format getDocumentTimestampFormat() {
        return TimestampFormatTrait.Format.DATE_TIME;
    }

    @Override
    protected String getDocumentContentType() {
        return null;
    }


    @Override
    protected void generateOperationDocumentSerializer(GenerationContext context, OperationShape operation) {
        Model model = context.getModel();
        HttpBindingIndex bindingIndex = model.getKnowledge(HttpBindingIndex.class);

        Set<MemberShape> documentBindings = bindingIndex.getRequestBindings(operation, HttpBinding.Location.DOCUMENT)
                .stream()
                .map(HttpBinding::getMember)
                .collect(Collectors.toSet());

        if (documentBindings.size() == 0) {
            return;
        }

        Shape inputShape = ProtocolUtils.expectInput(context.getModel(), operation);
        inputShape.accept(new XmlShapeSerVisitor(context, documentBindings::contains));

    }

    @Override
    protected void writeMiddlewareDocumentSerializerDelegator(
            GenerationContext context,
            OperationShape operation,
            GoStackStepMiddlewareGenerator generator
    ) {
        GoWriter writer = context.getWriter();
        writer.addUseImports(SmithyGoDependency.SMITHY);
        writer.addUseImports(SmithyGoDependency.SMITHY_XML);

        writer.write("restEncoder.SetHeader(\"Content-Type\").String($S)", getDocumentContentType());
        writer.write("");

        Shape inputShape = ProtocolUtils.expectInput(context.getModel(), operation);
        String functionName = ProtocolGenerator.getDocumentSerializerFunctionName(inputShape, getProtocolName());

        writer.addUseImports(SmithyGoDependency.SMITHY_XML);
        writer.write("xmlEncoder := smithyxml.NewEncoder()");

        writer.openBlock("if err := $L(input, xmlEncoder.Value.RootElement().Key($S)); err != nil {", "}",
                functionName, inputShape.getId().getName(), () -> {
            writer.write("return out, metadata, &smithy.SerializationError{Err: err}");
        });
        writer.write("");

        writer.addUseImports(SmithyGoDependency.BYTES);
        writer.openBlock("if request, err = request.SetStream(bytes.NewReader(xmlEncoder.Bytes())); err != nil {", "}",
                () -> {
                    writer.write("return out, metadata, &smithy.SerializationError{Err: err}");
                });
    }


    @Override
    protected void writeMiddlewarePayloadSerializerDelegator(
            GenerationContext context,
            OperationShape operation,
            MemberShape memberShape,
            GoStackStepMiddlewareGenerator generator
    ) {
        GoWriter writer = context.getWriter();
        Model model = context.getModel();
        Shape payloadShape = model.expectShape(memberShape.getTarget());

        writeSafeMemberAccessor(context, memberShape, "input", s -> {
            writer.openBlock("if !restEncoder.HasHeader(\"Content-Type\") {", "}", () -> {
                writer.write("restEncoder.SetHeader(\"Content-Type\").String($S)", getPayloadShapeMediaType(payloadShape));
            });
            writer.write("");

            if (payloadShape.hasTrait(StreamingTrait.class)) {
                writer.write("payload := $L", s);

            } else if (payloadShape.isBlobShape()) {
                writer.addUseImports(SmithyGoDependency.BYTES);
                writer.write("payload := bytes.NewReader($L)", s);

            } else if (payloadShape.isStringShape()) {
                writer.addUseImports(SmithyGoDependency.STRINGS);
                writer.write("payload := strings.NewReader(*$L)", s);

            } else {
                String functionName = ProtocolGenerator.getDocumentSerializerFunctionName(payloadShape,
                        getProtocolName());
                writer.addUseImports(SmithyGoDependency.SMITHY_XML);
                writer.write("xmlEncoder := smithyxml.NewEncoder()");
                writer.openBlock("if err := $L($L, xmlEncoder.Value); err != nil {", "}", functionName,
                        s, () -> {
                            writer.write("return out, metadata, &smithy.SerializationError{Err: err}");
                        });
                writer.write("payload := bytes.NewReader(xmlEncoder.Bytes())");
            }

            writer.openBlock("if request, err = request.SetStream(payload); err != nil {", "}",
                    () -> {
                        writer.write("return out, metadata, &smithy.SerializationError{Err: err}");
                    });
        });
    }

    @Override
    protected void generateDocumentBodyShapeSerializers(
            GenerationContext context, Set<Shape> shapes
    ) {
        XmlShapeSerVisitor visitor = new XmlShapeSerVisitor(context);
        shapes.forEach(shape -> shape.accept(visitor));
    }

    /**
     * Returns the MediaType for the payload shape derived from the MediaTypeTrait, shape type, or document content type.
     *
     * @param payloadShape shape bound to the payload.
     * @return string for media type.
     */
    private String getPayloadShapeMediaType(Shape payloadShape) {
        Optional<MediaTypeTrait> mediaTypeTrait = payloadShape.getTrait(MediaTypeTrait.class);

        if (mediaTypeTrait.isPresent()) {
            return mediaTypeTrait.get().getValue();
        }

        if (payloadShape.isBlobShape()) {
            return "application/octet-stream";
        }

        if (payloadShape.isStringShape()) {
            return "text/plain";
        }

        return getDocumentContentType();
    }



    /*========================Deserializers==================================*/
    @Override
    protected void writeMiddlewareDocumentDeserializerDelegator(
            GenerationContext context,
            OperationShape operation,
            GoStackStepMiddlewareGenerator generator
    ) {

    }

    @Override
    protected void generateOperationDocumentDeserializer(
            GenerationContext context, OperationShape operation
    ) {

    }

    @Override
    protected void generateDocumentBodyShapeDeserializers(
            GenerationContext context, Set<Shape> shapes
    ) {

    }

    @Override
    public ShapeId getProtocol() {
        return null;
    }

    @Override
    protected void writeErrorMessageCodeDeserializer(GenerationContext context) {
        // TODO: change this
        context.getWriter().writeDocs("TODO: implement error message / code deser");
        context.getWriter().addUseImports(SmithyGoDependency.FMT);
        context.getWriter().write("fmt.Println(errorBody)");
    }

    @Override
    protected void deserializeError(GenerationContext context, StructureShape shape) {
        GoWriter writer = context.getWriter();

        writer.write("return &smithy.DeserializationError{"
                + "Err: fmt.Errorf(\"TODO: Implement error deserializerfor %v\", errorBody)}");
    }
}
