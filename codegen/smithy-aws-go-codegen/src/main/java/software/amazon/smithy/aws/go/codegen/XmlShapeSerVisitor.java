package software.amazon.smithy.aws.go.codegen;

import static software.amazon.smithy.go.codegen.integration.ProtocolUtils.writeSafeMemberAccessor;

import java.util.Collections;
import java.util.Map;
import java.util.Optional;
import java.util.Set;
import java.util.TreeSet;
import java.util.function.Predicate;
import java.util.logging.Logger;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.integration.DocumentShapeSerVisitor;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator.GenerationContext;
import software.amazon.smithy.model.shapes.CollectionShape;
import software.amazon.smithy.model.shapes.DocumentShape;
import software.amazon.smithy.model.shapes.MapShape;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.model.shapes.UnionShape;
import software.amazon.smithy.model.traits.EnumTrait;
import software.amazon.smithy.model.traits.TimestampFormatTrait;
import software.amazon.smithy.model.traits.XmlNameTrait;
import software.amazon.smithy.utils.FunctionalUtils;

final class XmlShapeSerVisitor extends DocumentShapeSerVisitor {
 private static final TimestampFormatTrait.Format DEFAULT_TIMESTAMP_FORMAT = TimestampFormatTrait.Format.DATE_TIME;
 private static final Logger LOGGER = Logger.getLogger(JsonShapeSerVisitor.class.getName());

 private final Predicate<MemberShape> memberFilter;

 public XmlShapeSerVisitor(
         GenerationContext context
 ) {
  this(context, FunctionalUtils.alwaysTrue());
 }

 public XmlShapeSerVisitor(GenerationContext context,  Predicate<MemberShape> memberFilter) {
  super(context);
  this.memberFilter = memberFilter;
 }

 private XmlMemberSerVisitor getMemberSerVisitor(MemberShape member, String source, String dest) {
  // Get the timestamp format to be used, defaulting to epoch seconds.
  TimestampFormatTrait.Format format = member.getMemberTrait(getContext().getModel(), TimestampFormatTrait.class)
          .map(TimestampFormatTrait::getFormat).orElse(DEFAULT_TIMESTAMP_FORMAT);
  return new XmlMemberSerVisitor(getContext(), source, dest, format);
 }

 @Override
 protected Map<String, String> getAdditionalSerArguments() {
  return Collections.singletonMap("value", "smithyxml.Value");
 }

 @Override
 protected void serializeCollection(GenerationContext context, CollectionShape shape) {
  GoWriter writer = context.getWriter();
//  Shape target = context.getModel().expectShape(shape.getMember().getTarget());
//
//  writer.write("array := value.Array()");
//  writer.write("defer array.Close()");
//  writer.insertTrailingNewline();
//
//  writer.openBlock("for i := range v {", "}", () -> {
//   writer.write("av := array.Member()");
//   // Null values in lists should be serialized as such. Enums can't be null, so we don't bother
//   // putting this in for their case.
//   if (!target.hasTrait(EnumTrait.class)) {
//    writer.openBlock("if vv := v[i]; vv == nil {", "}", () -> {
//     writer.write("av.Null()");
//     writer.write("continue");
//    });
//   }
//
//   target.accept(getMemberSerVisitor(shape.getMember(), "v[i]", "av"));
//  });

  writer.write("return nil");
 }

 @Override
 protected void serializeDocument(GenerationContext context, DocumentShape shape) {
  // TODO: implement document serialization
  LOGGER.warning("Document type is currently unsupported for XML serialization.");
  context.getWriter().writeDocs("TODO: implement document serialization.");
  context.getWriter().write("return nil");
 }

 @Override
 protected void serializeMap(GenerationContext context, MapShape shape) {
  GoWriter writer = context.getWriter();
//  Shape target = context.getModel().expectShape(shape.getValue().getTarget());
//
//  writer.write("m := value.Map()");
//  writer.write("defer m.Close()");
//  writer.insertTrailingNewline();
//
//  writer.openBlock("for key := range v {", "}", () -> {
//
//   writer.write("entry := m.Entry()");
//   writer.write("defer entry.Close()");
//   writer.insertTrailingNewline();
//
//   writer.write("mv := entry.Key(key, nil)");
//
//   // Null values in maps should be serialized as such. Enums can't be null, so we don't bother
//   // putting this in for their case.
//   if (!target.hasTrait(EnumTrait.class)) {
//    writer.openBlock("if vv := v[key]; vv == nil {", "}", () -> {
//     writer.write("mv.Null()");
//     writer.write("continue");
//    });
//   }
//
//   target.accept(getMemberSerVisitor(shape.getValue(), "v[key]", "mv"));
//  });

  writer.write("return nil");
 }

 @Override
 protected void serializeStructure(GenerationContext context, StructureShape shape) {
  GoWriter writer = context.getWriter();
//
//  writer.write("ne := value.NestedElement()");
//  writer.write("defer ne.Close()");
//  writer.insertTrailingNewline();
//
//  // Use a TreeSet to sort the members.
//  Set<MemberShape> members = new TreeSet<>(shape.getAllMembers().values());
//  for (MemberShape member : members) {
//   if (!memberFilter.test(member)) {
//    continue;
//   }
//   Shape target = context.getModel().expectShape(member.getTarget());
//   String serializedMemberName = getSerializedMemberName(member);
//   writeSafeMemberAccessor(context, member, "v", (operand) -> {
//    writer.write("ek := ne.Key($S, nil)", serializedMemberName);
//    target.accept(getMemberSerVisitor(member, operand, "ek"));
//   });
//   writer.insertTrailingNewline();
//  }

  writer.write("return nil");
 }

 @Override
 protected void serializeUnion(
         GenerationContext context, UnionShape shape
 ) {
  // TODO: implement document serialization
  LOGGER.warning("Union type is currently unsupported for XML serialization.");
  context.getWriter().writeDocs("TODO: implement union serialization.");
  context.getWriter().write("return nil");
 }

 private String getSerializedMemberName(MemberShape memberShape) {
  Optional<XmlNameTrait> xmlNameTrait = memberShape.getTrait(XmlNameTrait.class);
  return xmlNameTrait.isPresent() ? xmlNameTrait.get().getValue() : memberShape.getMemberName();
 }
}
