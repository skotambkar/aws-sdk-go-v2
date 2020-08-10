package software.amazon.smithy.aws.go.codegen;

import java.util.Collection;
import java.util.Collections;
import java.util.Map;
import java.util.Optional;
import java.util.Set;
import java.util.TreeSet;
import java.util.function.Predicate;
import java.util.logging.Logger;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.integration.DocumentShapeDeserVisitor;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator.GenerationContext;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.CollectionShape;
import software.amazon.smithy.model.shapes.DocumentShape;
import software.amazon.smithy.model.shapes.MapShape;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.shapes.ShapeVisitor;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.model.shapes.UnionShape;
import software.amazon.smithy.model.traits.TimestampFormatTrait;
import software.amazon.smithy.model.traits.XmlAttributeTrait;
import software.amazon.smithy.model.traits.XmlFlattenedTrait;
import software.amazon.smithy.model.traits.XmlNameTrait;
import software.amazon.smithy.utils.FunctionalUtils;

/**
 * Visitor to generate deserialization functions for shapes in XML protocol
 * document bodies.
 *
 * This class handles function body generation for all types expected by the
 * {@code DocumentShapeDeserVisitor}. No other shape type serialization is overwritten.
 *
 * Timestamps are serialized to {@link TimestampFormatTrait.Format}.DATE_TIME by default.
 */
public class XmlShapeDeserVisitor extends DocumentShapeDeserVisitor {
    private static final TimestampFormatTrait.Format DEFAULT_TIMESTAMP_FORMAT = TimestampFormatTrait.Format.DATE_TIME;
    private static final Logger LOGGER = Logger.getLogger(XmlShapeDeserVisitor.class.getName());

    private final Predicate<MemberShape> memberFilter;

    /**
     * @param context The generation context.
     */
    public XmlShapeDeserVisitor(GenerationContext context) {
        this(context, FunctionalUtils.alwaysTrue());
    }

    /**
     * @param context The generation context.
     * @param memberFilter A filter that is applied to structure members. This is useful for
     *     members that won't be in the body.
     */
    public XmlShapeDeserVisitor(GenerationContext context, Predicate<MemberShape> memberFilter) {
        super(context);
        this.memberFilter = memberFilter;
    }

    @Override
    protected Map<String, String> getAdditionalArguments() {
        return Collections.singletonMap("decoder", "*smithydecoding.XMLNodeDecoder");
    }

    private XmlMemberDeserVisitor getMemberDeserVisitor(MemberShape member, String dataDest, Predicate<Shape> filter) {
        // Get the timestamp format to be used, defaulting to epoch seconds.
        TimestampFormatTrait.Format format = member.getMemberTrait(getContext().getModel(), TimestampFormatTrait.class)
                .map(TimestampFormatTrait::getFormat).orElse(DEFAULT_TIMESTAMP_FORMAT);

        // Check if FlattenedTrait  is applied
        Boolean isFlattened = member.hasTrait(XmlFlattenedTrait.ID);

        return new XmlMemberDeserVisitor(getContext(), dataDest, format, filter, isFlattened);
    }

    @Override
    protected void deserializeCollection(GenerationContext context, CollectionShape shape) {
        GoWriter writer = context.getWriter();
        SymbolProvider symbolProvider = context.getSymbolProvider();
        Symbol symbol = symbolProvider.toSymbol(shape);

        // Initialize the value now that the start stub has verified that there's something there.
        writer.write("var sv $P", symbol);
        writer.openBlock("if *v == nil {", "", () -> {
            writer.write("sv = make($P, 0)", symbol);
            writer.openBlock("} else {", "}", () -> {
                writer.write("sv = *v");
            });
        });

        // Iterate through the decoder. The member visitor will handle popping tokens.
        writer.openBlock("for {", "}", () -> {
            writer.write("t, done, err := decoder.Token()");
            writer.write("if err != nil { return err }");
            writer.write("if done { break }");


            // non-flattened maps
            MemberShape member = shape.getMember();
            Shape target = context.getModel().expectShape(member.getTarget());
            String memberName = symbolProvider.toMemberName(member);
            String serializedMemberName = getSerializedMemberName(member);
            writer.openBlock("switch t.Name.Local {", "}", () -> {
                writer.openBlock("case $S:", "", serializedMemberName, () -> {
                    writer.write("var col $P", context.getSymbolProvider().toSymbol(target));
                    target.accept(getMemberDeserVisitor(member, "col", FunctionalUtils.alwaysTrue()));
                    writer.write("sv = append(sv, col)");
                });

                writer.openBlock("default:", "", () -> {
                    writer.write("// Do nothing and ignore");
                });
            });
        });
        writer.write("*v = sv");
        writer.write("return nil");
    }

    public void generateFlattenedCollection(GenerationContext context, CollectionShape shape) {
        GoWriter writer = context.getWriter();
        SymbolProvider symbolProvider = context.getSymbolProvider();
        Symbol symbol = symbolProvider.toSymbol(shape);

        MemberShape member = shape.getMember();
        Symbol memberSymbol = symbolProvider.toSymbol(member);

        Shape target  = context.getModel().expectShape(member.getTarget());
        Symbol targetSymbol = symbolProvider.toSymbol(target);

        writer.openBlock("func $L(v *$P, decoder *smithydecoding.XMLNodeDecoder) error {", "}",
                getUnwrappedMapDelegateFunctionName(context, shape), symbol, () -> {
                    // Initialize the value now that the start stub has verified that there's something there.
                    writer.write("var sv $P", symbol);
                    writer.openBlock("if *v == nil {", "", () -> {
                        writer.write("sv = make($P, 0)", symbol);
                        writer.openBlock("} else {", "}", () -> {
                            writer.write("sv = *v");
                        });
                    });

                    writer.openBlock("for {","}", () -> {
                        writer.write("var mv $P", memberSymbol);
                        writer.write("t := decoder.StartEl");
                        writer.write("_ = t");
                        String dest = "mv";
                        target.accept(getMemberDeserVisitor(member, dest, FunctionalUtils.alwaysTrue()));
                        writer.write("sv = append(sv, mv)");
                    });

                    writer.write("*v = sv");
                    writer.write("return nil");
                });
    }

    @Override
    protected void deserializeMap(GenerationContext context, MapShape shape) {
        GoWriter writer = context.getWriter();
        Model model = context.getModel();
        SymbolProvider symbolProvider = context.getSymbolProvider();
        Symbol symbol = symbolProvider.toSymbol(shape);


        // Initialize the value now that the start stub has verified that there's something there.
        writer.write("var sv $P", symbol);
        writer.openBlock("if *v == nil {", "", () -> {
            writer.write("sv = make($P, 0)", symbol);
            writer.openBlock("} else {", "}", () -> {
                writer.write("sv = *v");
            });
        });

        // Iterate through the decoder. The member visitor will handle popping tokens.
        writer.openBlock("for {", "}", () -> {
            writer.write("t, done, err := decoder.Token()");
            writer.write("if err != nil { return err }");
            writer.write("if done { break }");

            // non-flattened maps
            writer.openBlock("switch t.Name.Local {", "}", () -> {
                writer.write("case \"entry\":");
                writer.write("entryDecoder := smithydecoding.NewXMLNodeDecoder(decoder.Decoder, t)");
                // delegate to unwrapped map deserializer function
                writer.openBlock("if err := $L(&sv, entryDecoder); err != nil {", "}",
                        getUnwrappedMapDelegateFunctionName(context, shape), () -> {
                    writer.write("return err");
                    });

                writer.openBlock("default:", "", () -> {
                    writer.write("// Do nothing and ignore");
                });
            });
        });
        writer.write("*v = sv");
        writer.write("return nil");
    }

    public void generateFlattenedMap(GenerationContext context, MapShape shape) {
        GoWriter writer = context.getWriter();
        SymbolProvider symbolProvider = context.getSymbolProvider();
        Symbol symbol = symbolProvider.toSymbol(shape);

        writer.openBlock("func $L(v *$P, decoder *smithydecoding.XMLNodeDecoder) error {", "}",
                getUnwrappedMapDelegateFunctionName(context, shape), symbol, () -> {
                    // Initialize the value now that the start stub has verified that there's something there.
                    writer.write("var sv $P", symbol);
                    writer.openBlock("if *v == nil {", "", () -> {
                        writer.write("sv = make($P, 0)", symbol);
                        writer.openBlock("} else {", "}", () -> {
                            writer.write("sv = *v");
                        });
                    });

                    writer.write("var ek $P", symbolProvider.toSymbol(shape.getKey()));
                    writer.write("var ev $P", symbolProvider.toSymbol(shape.getValue()));
                    writer.insertTrailingNewline();

                    // Iterate through the decoder. The member visitor will handle popping tokens.
                    writer.openBlock("for {", "}", () -> {
                        writer.write("t, done, err := decoder.Token()");
                        writer.write("if err != nil { return err }");
                        writer.openBlock("if done {", "}", () -> {
                            // set the key value pair in map
                            writer.write("sv[*ek] = ev");
                            writer.write("break");
                        });

                        writer.openBlock("switch t.Name.Local {", "}", () -> {
                            MemberShape keyShape = shape.getKey();
                            writer.openBlock("case $S:", "", getSerializedMemberName(keyShape), () -> {
                                String dest = "ek";
                                context.getModel().expectShape(keyShape.getTarget()).accept(getMemberDeserVisitor(keyShape, dest,((shape1)-> !shape1.hasTrait(XmlAttributeTrait.ID))));
                            });

                            MemberShape valueShape = shape.getValue();
                            writer.openBlock("case $S:", "", getSerializedMemberName(valueShape), () -> {
                                String dest = "ev";
                                context.getModel().expectShape(valueShape.getTarget()).accept(getMemberDeserVisitor(valueShape, dest,((shape1)-> !shape1.hasTrait(XmlAttributeTrait.ID))));
                            });

                            writer.openBlock("default:", "", () -> {
                                writer.write("// Do nothing and ignore");
                            });
                        });
                    });
                    writer.write("*v = sv");
                    writer.write("return nil");
        });
    }

    private String getUnwrappedMapDelegateFunctionName(GenerationContext context, Shape shape) {
        return ProtocolGenerator.getDocumentDeserializerFunctionName(shape, context.getProtocolName()) + "Unwrapped";
    }

    // Generate the unwrapped Map method

    protected void deserializeUnwrappedMap(GenerationContext context, MapShape shape) {
        GoWriter writer = context.getWriter();
        SymbolProvider symbolProvider = context.getSymbolProvider();
        Symbol symbol = symbolProvider.toSymbol(shape);

        writer.write("return nill");
    }

    private String getUnWrappedMapFunctionName(GenerationContext context, Shape shape) {
        return ProtocolGenerator.getDocumentDeserializerFunctionName(shape, context.getProtocolName()) + "Unwrapped";
    }



    @Override
    protected void deserializeStructure(GenerationContext context, StructureShape shape) {
        GoWriter writer = context.getWriter();
        SymbolProvider symbolProvider = context.getSymbolProvider();
        Symbol symbol = symbolProvider.toSymbol(shape);

//        writeXMLTokenizerStartStub(writer, shape);

        // Initialize the value now that the start stub has verified that there's something there.
        writer.write("var sv $P", symbol);
        writer.openBlock("if *v == nil {", "", () -> {
            writer.write("sv = &$T{}", symbol);
            writer.openBlock("} else {", "}", () -> {
                writer.write("sv = *v");
            });
        });

        // TODO: Refactor this to be more inline and DRY when handling xml attributes
        Predicate<Shape> hasXmlAttributeTraitMember = shape1 -> {
            Collection<MemberShape> members = shape1.members();
            for (MemberShape member: members) {
                if (member.hasTrait(XmlAttributeTrait.ID)) {
                    return true;
                }
            }
            return false;
        };

        if (hasXmlAttributeTraitMember.test(shape)) {
            writer.openBlock("for _, attr := range decoder.StartEl.Attr {", "}", () -> {
                writer.write("val := []byte(attr.Value)");
                writer.openBlock("switch attr.Name.Local {", "}", () -> {
                    Set<MemberShape> members = new TreeSet<>(shape.members());
                    for (MemberShape member : members) {
                        // check if member is not a document binding or does not have a xmlAttribute trait
                        if (!memberFilter.test(member) || !member.hasTrait(XmlAttributeTrait.ID)) {
                            continue;
                        }

                        String memberName = symbolProvider.toMemberName(member);
                        String serializedMemberName = getSerializedMemberName(member);
                        writer.openBlock("case $S:", "", serializedMemberName, () -> {
                            String dest = "sv." + memberName;
                            context.getModel().expectShape(member.getTarget()).accept(
                                    getMemberDeserVisitor(member, dest, ((shape1) -> shape1.hasTrait(XmlAttributeTrait.ID))));
                        });
                    }
                });
            });
        }

        // Iterate through the decoder. The member visitor will handle popping tokens.
        writer.openBlock("for {", "}", () -> {
            writer.write("t, done, err := decoder.Token()");
            writer.write("if err != nil { return err }");
            writer.write("if done { break }");

            writer.openBlock("switch t.Name.Local {", "}", () -> {
                Set<MemberShape> members = new TreeSet<>(shape.members());
                for (MemberShape member : members) {
                    // check if member is not a document binding or has a xmlAttribute trait
                    if (!memberFilter.test(member) || member.hasTrait(XmlAttributeTrait.ID)) {
                        continue;
                    }

                    String memberName = symbolProvider.toMemberName(member);
                    String serializedMemberName = getSerializedMemberName(member);
                    writer.openBlock("case $S:", "", serializedMemberName, () -> {
                        String dest = "sv." + memberName;
                        context.getModel().expectShape(member.getTarget()).accept(getMemberDeserVisitor(member, dest,((shape1)-> !shape1.hasTrait(XmlAttributeTrait.ID))));
                    });
                }

                writer.openBlock("default:", "", () -> {
                   writer.write("// Do nothing and ignore");
                });
            });
        });

//        writeXMLTokenizerEndStub(writer, shape);
        writer.write("*v = sv");
        writer.write("return nil");
    }

    private String getSerializedMemberName(MemberShape memberShape) {
        Optional<XmlNameTrait> xmlNameTrait = memberShape.getTrait(XmlNameTrait.class);
        return xmlNameTrait.isPresent() ? xmlNameTrait.get().getValue() : memberShape.getMemberName();
    }

    @Override
    protected void deserializeDocument(GenerationContext context, DocumentShape shape) {
        GoWriter writer = context.getWriter();
        LOGGER.warning("The document type is unsupported for XML protocols.");
        writer.addUseImports(SmithyGoDependency.SMITHY);
        writer.write("return &smithy.DeserializationError{Err: fmt.Errorf("
                + "\"Document type is unsupported for XML protocols.\")}");
    }

    @Override
    protected void deserializeUnion(GenerationContext context, UnionShape shape) {
        GoWriter writer = context.getWriter();
        // TODO: implement union deserialization
        // The tricky part is going to be accumulating bytes for unknown members.
        LOGGER.warning("Union type is currently unsupported for XML deserialization.");
        context.getWriter().writeDocs("TODO: implement union serialization.");
        writer.write("return nil");
    }

    @Override
    protected Void getDefault(Shape shape) {
        if (shape.isMapShape()) {
            generateDeserFunction(shape , getUnWrappedMapFunctionName(super.getContext(), shape), (c, s) -> deserializeUnwrappedMap(c, s.asMapShape().get()));
        }
        return null;
    }

    /*=================================== Helpers ===========================================*/

    /**
     * Writes out a stub to initialize decoding.
     *
     * @param writer The GoWriter to use.
     * @param shape The shape the stub is intended to start arsing.
     */
    private void writeXMLTokenizerStartStub(GoWriter writer, Shape shape) {
        writer.addUseImports(SmithyGoDependency.XML);
        writer.write("startToken, err := decoder.Token()");
        writer.write("if err != nil { return err }");
        writer.write("if startToken == nil { return nil }");
        writer.openBlock("if t, ok := startToken.(xml.StartElement); !ok {","}", () -> {
            writer.addUseImports(SmithyGoDependency.FMT);
            writer.write("return fmt.Errorf(\"expected start token, got %v\", startToken)");
        });
        writer.insertTrailingNewline();
    }

    /**
     * Writes out a stub to finalize decoding.
     *
     * @param writer The GoWriter to use.
     * @param shape The shape the stub is intended to finalize parsing for.
     */
    private void writeXMLTokenizerEndStub(GoWriter writer, Shape shape) {
        writer.addUseImports(SmithyGoDependency.XML);
        writer.write("endToken, err := decoder.Token()");
        writer.write("if err != nil { return err }");
        writer.openBlock("if t, ok := endToken.(xml.EndElement); !ok {", "}", () -> {
            writer.write("return fmt.Errorf(\"expected end element, got %v\",endToken)");
        });
        writer.insertTrailingNewline();
    }
}
