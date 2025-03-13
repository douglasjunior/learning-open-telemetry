using OpenTelemetry.Exporter;
using OpenTelemetry.Metrics;
using OpenTelemetry.Trace;
using OpenTelemetry.Resources;
using OpenTelemetry;

namespace LearningOtelDotnet.Configs;

public static class OpenTelemetryConfig
{
  public static void AddOpenTelemetryInstrumentation(this IServiceCollection services, ILoggingBuilder logging, ConfigurationManager configuration)
  {
    var otelProtocol = (OtlpExportProtocol)Enum.Parse(
      typeof(OtlpExportProtocol),
      // Grpc or HttpProtobuf
      configuration.GetValue("Otlp:Protocol", defaultValue: "Grpc"),
      true
    );
    var otelEndpoint = new Uri(configuration.GetValue("Otlp:Endpoint", defaultValue: "http://localhost:4317"));

    logging.AddOpenTelemetry();
    services.AddOpenTelemetry()
        .WithMetrics(metrics => metrics.AddAspNetCoreInstrumentation()
            .AddHttpClientInstrumentation()
            .AddSqlClientInstrumentation())
        .WithTracing(tracing => tracing.AddAspNetCoreInstrumentation()
            .AddHttpClientInstrumentation()
            .AddEntityFrameworkCoreInstrumentation(null, options => options.SetDbStatementForText = true))
        .ConfigureResource(resourceBuilder => resourceBuilder.AddService("LearningOtelDotnet"))
        .UseOtlpExporter(
            otelProtocol,
            otelEndpoint
        );
  }
}
