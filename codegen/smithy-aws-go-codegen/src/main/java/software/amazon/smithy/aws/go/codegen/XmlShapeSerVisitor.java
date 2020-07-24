package software.amazon.smithy.aws.go.codegen;

import static software.amazon.smithy.go.codegen.integration.ProtocolUtils.writeSafeMemberAccessor;

import java.util.Collections;
import java.util.Map;
import java.util.Optional;
import java.util.Set;
import java.util.TreeSet;
import java.util.function.Predicate;
import java.util.logging.Logger;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.go.codegen.GoDependency;
import software.amazon.smithy.go.codegen.GoSettings;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.SymbolUtils;
import software.amazon.smithy.go.codegen.SyntheticClone;
import software.amazon.smithy.go.codegen.integration.DocumentShapeSerVisitor;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator.GenerationContext;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.knowledge.TopDownIndex;
import software.amazon.smithy.model.shapes.CollectionShape;
import software.amazon.smithy.model.shapes.DocumentShape;
import software.amazon.smithy.model.shapes.MapShape;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.OperationShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.shapes.ShapeId;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.model.shapes.UnionShape;
import software.amazon.smithy.model.traits.EnumTrait;
import software.amazon.smithy.model.traits.TimestampFormatTrait;
import software.amazon.smithy.model.traits.XmlAttributeTrait;
import software.amazon.smithy.model.traits.XmlFlattenedTrait;
import software.amazon.smithy.model.traits.XmlNameTrait;
import software.amazon.smithy.model.traits.XmlNamespaceTrait;
import software.amazon.smithy.utils.CaseUtils;
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

 // getSerializedShapeName retrieves the shape name to be used for serialization.
 // The xml name is returned if it is defined.
 private static String getSerializedShapeName(GenerationContext context, Shape shape) {
  SymbolProvider symbolProvider = context.getSymbolProvider();
  Symbol shapeSymbol  =  symbolProvider.toSymbol(shape);
  String shapeName = shapeSymbol.getName();

  // check if synthetic cloned shape
  Optional<SyntheticClone> clone = shape.getTrait(SyntheticClone.class);
  if (clone.isPresent()){
   SyntheticClone cl = clone.get();
   shapeName = cl.getArchetype().getName();
  }

  // check if shape is member shape
  Optional<MemberShape> member = shape.asMemberShape();
  if (member.isPresent()) {
   MemberShape memberShape = member.get();
   return getSerializedMemberName(memberShape);
  }

  Optional<XmlNameTrait> xmlNameTrait = shape.getTrait(XmlNameTrait.class);
  return xmlNameTrait.isPresent() ? xmlNameTrait.get().getValue() : shapeName;
 }

 // TODO: Add documentation and move this
 public static void generateXMLStartElementStub(
         GenerationContext context, Shape shape, String dst, String inputSrc
 ) {
  SymbolProvider symbolProvider = context.getSymbolProvider();
  GoWriter writer = context.getWriter();
  String attrName = dst +"Attr";
  writer.write("$L := []smithyxml.Attr{}", attrName);

  Optional<XmlNamespaceTrait> xmlNamespaceTrait = shape.getTrait(XmlNamespaceTrait.class);
  if (xmlNamespaceTrait.isPresent()) {
   XmlNamespaceTrait ns = xmlNamespaceTrait.get();
   writer.write("$L = append($L, smithyxml.NewNamespaceAttribute($S, $S))",
           attrName, attrName,
           ns.getPrefix().isPresent()? ns.getPrefix().get():"", ns.getUri()
   );
  }

  // Traverse member shapes to get attributes
  shape.members().stream().forEach(memberShape -> {
   Optional<XmlAttributeTrait> xmlAttributeTrait = memberShape.getTrait(XmlAttributeTrait.class);
   if (xmlAttributeTrait.isPresent()) {
    writeSafeMemberAccessor(context, memberShape, inputSrc, (operand)->{
     writer.write("$L = append($L, smithyxml.NewAttribute($S, string(*$L)))",
             attrName, attrName,
             getSerializedMemberName(memberShape),operand);
    });
   }
  });

  writer.openBlock("$L := smithyxml.StartElement{ ", "}", dst, () -> {
   writer.openBlock("Name:smithyxml.Name{","},", () -> {
    writer.write("Local: $S,", getSerializedShapeName(context,shape));
   });
   writer.write("Attr : $L,", attrName);
  });
 }

 @Override
 protected Map<String, String> getAdditionalSerArguments() {
  return Collections.singletonMap("value", "smithyxml.Value");
 }

 @Override
 protected void serializeCollection(GenerationContext context, CollectionShape shape) {
  GoWriter writer = context.getWriter();
  Shape target = context.getModel().expectShape(shape.getMember().getTarget());
  MemberShape member = shape.getMember();
  writer.write("var array *smithyxml.Array");

//   if (member.hasTrait(XmlNameTrait.class) || member.hasTrait(XmlNamespaceTrait.class)) {
//    writer.openBlock("if value.IsFlattened() {", "} else {", () ->{
//     writer.write("array = value.FlattenedArray()");
//    });
//    generateXMLStartElementStub(context, member, "customMemberName", "v");
//    writer.write("array = value.ArrayWithCustomName(customMemberName)");
//   } else{
//    writer.openBlock("if value.IsFlattened() {", "} else {", () ->{
//     writer.write("array = value.FlattenedArray()");
//    });
//    writer.write("array = value.Array()");
//   }

  // TODO: may be close it in the same function it came from.
  writer.openBlock("if !value.IsFlattened() {", "}", () -> {
   writer.write("defer value.Close()");
  });

  if (member.hasTrait(XmlNameTrait.class) || member.hasTrait(XmlNamespaceTrait.class)) {
   generateXMLStartElementStub(context, member, "customMemberName", "v");
   writer.write("array = value.ArrayWithCustomName(customMemberName)");
  } else {
   writer.write("array = value.Array()");
  }

  writer.insertTrailingNewline();

  writer.openBlock("for i := range v {", "}", () -> {
           // Null values in lists should be serialized as such. Enums can't be null, so we don't bother
           // putting this in for their case.
           if (!target.hasTrait(EnumTrait.class)) {
            writer.openBlock("if vv := v[i]; vv == nil {", "}", () -> {
             writer.write("am := array.Member()");
             writer.write("am.Close()");
             writer.write("continue");
            });
           }

           // Nested Array should use Collection member
           // TODO: check if this should include maps
//   if (target.isSetShape()|| target.isListShape()) {
//      writer.write("am := array.CollectionMember()");
//   } else {
           writer.write("am := array.Member()");
//          }
   target.accept(getMemberSerVisitor(shape.getMember(), "v[i]", "am"));
  });

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
  Shape targetKey = context.getModel().expectShape(shape.getKey().getTarget());
  Shape targetValue = context.getModel().expectShape(shape.getValue().getTarget());

//  writer.write("var m *smithyxml.Map");
//  writer.openBlock("if value.IsFlattened() { ", "} else {", () -> {
//   writer.write("m = value.FlattenedMap()");
//  });
//  writer.write("m = value.Map()");
//  writer.write("defer m.Close()}");

//  writer.write("var m *smithyxml.Map");
//  writer.openBlock("if value.IsFlattened() {", "} else {", () ->{
//   writer.write("m = value.FlattenedMap()");
//  });


  writer.openBlock("if !value.IsFlattened() {", "}", () -> {
   writer.write("defer value.Close()");
  });

  writer.write("m := value.Map()");

//  if (shape.hasTrait(XmlFlattenedTrait.class)
//          || shape.getMemberTrait(context.getModel(), XmlFlattenedTrait.class).isPresent()) {
//   writer.write("m := value.FlattenedMap()");
//  } else {
//   writer.write("m := value.Map()");
//   writer.write("defer m.Close()");
//  }
  writer.insertTrailingNewline();

  writer.openBlock("for key := range v {", "}", () -> {
   writer.write("entry := m.Entry()");
   writer.insertTrailingNewline();

   // Null values in maps should be serialized as such. Enums can't be null, so we don't bother
   // putting this in for their case.
   if (!targetValue.hasTrait(EnumTrait.class)) {
    writer.openBlock("if vv := v[key]; vv == nil {", "}", () -> {
     writer.write("entry.Close()");
     writer.write("continue");
    });
   }

   generateXMLStartElementStub(context, shape.getKey(), "keyElement", "v");
   targetKey.accept(getMemberSerVisitor(shape.getKey(), "&key", "entry.MemberElement(keyElement)"));


   String dest = "entry.MemberElement(valueElement)";
   // TODO should this be target?
   if (shape.getValue().hasTrait(XmlFlattenedTrait.class)) {
    dest = "entry.FlattenedElement(valueElement)";
   }

   generateXMLStartElementStub(context, shape.getValue(), "valueElement", "v");
   targetValue.accept(getMemberSerVisitor(shape.getValue(), "v[key]", dest));

   writer.write("entry.Close()");
  });
  writer.write("return nil");
 }

 @Override
 protected void serializeStructure(GenerationContext context, StructureShape shape) {
  GoWriter writer = context.getWriter();
  writer.write("defer value.Close()");
  writer.insertTrailingNewline();

  // Use a tree sort to sort the members.
  Set<MemberShape> members = new TreeSet<>(shape.getAllMembers().values());
  for (MemberShape member : members) {
   if (!memberFilter.test(member)) {
    continue;
   }

   Shape target = context.getModel().expectShape(member.getTarget());
   String serializedMemberName = getSerializedMemberName(member);

   // generate start element
   writer.addUseImports(SmithyGoDependency.SMITHY_XML);
   writeSafeMemberAccessor(context, member, "v", (operand) -> {
    // if member shape is a collection type use value.Collection element
    // else use value.Member element
    generateXMLStartElementStub(context, member, "root", "v");

    // TODO check if this trait is on target or member
    if (member.hasTrait(XmlFlattenedTrait.class)) {
     writer.write("el := value.FlattenedElement($L)", "root");
    } else {
     writer.write("el := value.MemberElement($L)", "root");
    }

//    // TODO: check if we have a util for these types
//    if (target.isMapShape() || target.isListShape() || target.isSetShape()) {
//     writer.write("el:= value.CollectionElement($L)", "root");
//     if (target.hasTrait(XmlFlattenedTrait.class) || member.hasTrait(XmlFlattenedTrait.class)) {
//      writer.write("el = el.SetFlattened()");
//     }
//    } else {
//     writer.write("el := value.MemberElement($L)", "root");
//    }

    // TODO: change this logic to be more sutle and user friendly
    // Check if shape is flattened


    target.accept(getMemberSerVisitor(member, operand, "el"));
   });

   writer.insertTrailingNewline();
  }

  writer.write("return nil");
 }
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

//  writer.write("return nil");
// }

 @Override
 protected void serializeUnion(
         GenerationContext context, UnionShape shape
 ) {
  // TODO: implement document serialization
  LOGGER.warning("Union type is currently unsupported for XML serialization.");
  context.getWriter().writeDocs("TODO: implement union serialization.");
  context.getWriter().write("return nil");
 }

 private static String getSerializedMemberName(MemberShape memberShape) {
  Optional<XmlNameTrait> xmlNameTrait = memberShape.getTrait(XmlNameTrait.class);
  return xmlNameTrait.isPresent() ? xmlNameTrait.get().getValue() : memberShape.getMemberName();
 }
}
