package software.amazon.smithy.aws.go.codegen;

import static software.amazon.smithy.go.codegen.integration.ProtocolUtils.writeSafeMemberAccessor;

import java.util.ArrayList;
import java.util.Collection;
import java.util.Optional;
import java.util.Set;
import java.util.stream.Collectors;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
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
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.model.traits.MediaTypeTrait;
import software.amazon.smithy.model.traits.StreamingTrait;
import software.amazon.smithy.model.traits.TimestampFormatTrait;
import software.amazon.smithy.model.traits.XmlAttributeTrait;
import software.amazon.smithy.model.traits.XmlNameTrait;
import software.amazon.smithy.model.traits.XmlNamespaceTrait;

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

    // getSerializedShapeName retrieves the shape name to be used for serialization.
    // The xml name is returned if it is defined.
    private String getSerializedShapeName(Shape shape) {
        Optional<XmlNameTrait> xmlNameTrait = shape.getTrait(XmlNameTrait.class);
        return xmlNameTrait.isPresent() ? xmlNameTrait.get().getValue() : shape.getId().getName();
    }

    private void generateXMLStartElementStub (GenerationContext context, GoWriter writer, Shape shape, String dst, String inputSrc) {
        SymbolProvider symbolProvider = context.getSymbolProvider();
        writer.write("attr := []smithyxml.Attr{}");

        Optional<XmlNamespaceTrait> xmlNamespaceTrait = shape.getTrait(XmlNamespaceTrait.class);
        if (xmlNamespaceTrait.isPresent()) {
            XmlNamespaceTrait ns = xmlNamespaceTrait.get();
            writer.write("attr = append(attr, smithyxml.NewNamespaceAttribute($S, $S))",
                    ns.getPrefix().isPresent()? ns.getPrefix().get():"", ns.getUri()
            );
        }

        // Traverse member shapes to get attributes
        shape.members().stream().forEach(memberShape -> {
            Optional<XmlAttributeTrait> xmlAttributeTrait = memberShape.getTrait(XmlAttributeTrait.class);
            if (xmlAttributeTrait.isPresent()){
                writeSafeMemberAccessor(context, memberShape, inputSrc, (operand)->{
                    writer.write("attr = append(attr, smithyxml.NewAttribute($S, string(*$L)))",
                            getSerializedShapeName(memberShape),operand);
                });
            }
        });

        writer.openBlock("$L = smithyxml.StartElement{ ", "}", dst, () -> {
            writer.openBlock("Name:smithyxml.Name{","},", () -> {
                writer.write("Local: $S,", getSerializedShapeName(shape));
            });
            writer.write("Attr : attr,");
        });
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
        writer.addUseImports(SmithyGoDependency.BYTES);
        writer.write("xmlEncoder := smithyxml.NewEncoder(bytes.NewBuffer(nil))");
        // TODO: fux thsi
//        writer.write("root := ", getXMLStartElement(input, inputShape));
        writer.write("var root smithyxml.StartElement");
        generateXMLStartElementStub(context, writer, inputShape, "root", "input");

        writer.openBlock("if err := $L(input, xmlEncoder.RootElement(root)); err != nil {", "}",
                functionName, () -> {
            writer.write("return out, metadata, &smithy.SerializationError{Err: err}");
        });
        writer.insertTrailingNewline();

        writer.openBlock("if request, err = request.SetStream(bytes.NewReader(xmlEncoder.Bytes())); " +
                        "err != nil {", "}", () -> {
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
                writer.addUseImports(SmithyGoDependency.BYTES);
                writer.write("xmlEncoder := smithyxml.NewEncoder(bytes.NewBuffer(nil))");

                // TODO: verify PayloadShape needs to be wrapped
                writer.write("var payloadRoot smithyxml.StartElement");
                generateXMLStartElementStub(context, writer, payloadShape, "payloadRoot", s);
                writer.openBlock("if err := $L($L, xmlEncoder.RootElement(payloadRoot)); err != nil {", "}", functionName,
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
    protected void writeErrorMessageCodeDeserializer(GenerationContext context) {
        context.getWriter().writeDocs("TODO: implement error message / code deser");
        context.getWriter().write("_= errorBody");
    }

    @Override
    protected void deserializeError(GenerationContext context, StructureShape shape) {
        GoWriter writer = context.getWriter();

        writer.write("return &smithy.DeserializationError{"
                + "Err: fmt.Errorf(\"TODO: Implement error deserializerfor %v\", errorBody)}");
    }
}
