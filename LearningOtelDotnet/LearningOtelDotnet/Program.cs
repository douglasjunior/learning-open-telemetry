using LearningOtelDotnet.Models;
using LearningOtelDotnet.Configs;
using LearningOtelDotnet.Client;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddDbContext<AppDbContext>();

builder.Services.AddHttpClients(builder.Configuration);
builder.Services.AddServices();
builder.Services.AddOpenTelemetryInstrumentation(builder.Logging, builder.Configuration);

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.MapControllers();

await app.RunAsync();
