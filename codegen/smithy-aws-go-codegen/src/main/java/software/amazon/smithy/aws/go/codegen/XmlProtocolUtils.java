package software.amazon.smithy.aws.go.codegen;

import software.amazon.smithy.aws.traits.protocols.RestXmlTrait;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator;

final class XmlProtocolUtils {
    private XmlProtocolUtils() {

    }

    /**
     * initializeXmlDecoder generates stub code to initialize xml decoder
     *
     * @param writer       the go writer used to write
     * @param bodyLocation the variable used to represent response body
     */
    public static void initializeXmlDecoder(GoWriter writer, String bodyLocation) {
        // Use a ring buffer and tee reader to help in pinpointing any deserialization errors.
        writer.addUseImports(SmithyGoDependency.SMITHY_IO);
        writer.write("buff := make([]byte, 1024)");
        writer.write("ringBuffer := smithyio.NewRingBuffer(buff)");
        writer.insertTrailingNewline();

        writer.addUseImports(SmithyGoDependency.IO);
        writer.addUseImports(SmithyGoDependency.XML);
        writer.addUseImports(SmithyGoDependency.SMITHY_DECODING);
        writer.write("body := io.TeeReader($L, ringBuffer)", bodyLocation);
        writer.write("rootDecoder := xml.NewDecoder(body)");
        writer.write("t, err := smithydecoding.FetchXmlRootElement(rootDecoder)");

        handleDecodeError(writer, "");
        writer.write("decoder := smithydecoding.NewXMLNodeDecoder(rootDecoder, t)");
        writer.insertTrailingNewline();
    }


    /**
     * handleDecodeError handles the xml deserialization error wrapping.
     *
     * @param writer the go writer used to write
     * @param returnExtras extra variables to be returned with the wrapped error statement
     */
    public static void handleDecodeError(GoWriter writer, String returnExtras) {
        writer.addUseImports(SmithyGoDependency.IO);
        writer.openBlock("if err != nil && err != io.EOF {", "}", () -> {
            writer.addUseImports(SmithyGoDependency.BYTES);
            writer.addUseImports(SmithyGoDependency.SMITHY);
            writer.write("var snapshot bytes.Buffer");
            writer.write("io.Copy(&snapshot, ringBuffer)");
            writer.openBlock("return $L&smithy.DeserializationError {", "}", returnExtras, () -> {
                writer.write("Err : fmt.Errorf(\"failed to decode response body, %w\", err),");
                writer.write("Snapshot: snapshot.Bytes(),");
            });
        }).write("");
    }

    /**
     * Generates code to retrieve error code or error message from the error response body
     * This method is used indirectly by generateErrorDispatcher to generate operation specific error handling functions
     *
     * @param context the generation context
     * @see <a href="https://awslabs.github.io/smithy/1.0/spec/aws/aws-restxml-protocol.html#operation-error-serialization">Rest-XML operation error serialization.</a>
     */
    public static void writeXmlErrorMessageCodeDeserializer(ProtocolGenerator.GenerationContext context) {
        GoWriter writer = context.getWriter();

        // Check if service uses isNoErrorWrapping setting
        boolean isNoErrorWrapping = context.getService().getTrait(RestXmlTrait.class).map(
                RestXmlTrait::isNoErrorWrapping).orElse(false);

        writer.write("errorCode, err := smithydecoding.GetXMLResponseErrorCode(errorBody, $L)", isNoErrorWrapping);
        writer.write("if err != nil { return err }");
        writer.insertTrailingNewline();

        writer.write("errorBody.Seek(0, io.SeekStart)");
        writer.insertTrailingNewline();
    }
}
