using System.Collections.Generic;
using Pulumi;
using Pulumi.Honeycomb.Resources;

return await Deployment.RunAsync(() =>
{
   // Add your resources here

   var dataset = new Dataset("my-pulumi-dataset", new DatasetArgs {
      Name = "my pulumi dataset",
      Description = "Created from pulumi",
      Expand_json_depth = 3
   });

   var column = new Column("my-pulumi-column2", new ColumnArgs {
      Name = "my-pulumi-column2",
      Dataset = dataset.Slug,
      Type = ColumnType.String,
      Description = "Created from pulumi",
      Hidden = false
   });

   // Export outputs here
   return new Dictionary<string, object?>
   {
      ["outputKey"] = "outputValue"
   };
});
