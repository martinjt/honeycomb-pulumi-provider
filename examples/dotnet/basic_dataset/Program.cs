using System.Collections.Generic;
using Pulumi;
using Pulumi.Honeycomb.Resources;

return await Deployment.RunAsync(() =>
{
   // Add your resources here

   var dataset = new Dataset("my-pulumi-dataset", new DatasetArgs {
      Name = "my-pulumi-dataset",
      Description = "Created from pulumi",
      Expand_json_depth = 3
   });

   // Export outputs here
   return new Dictionary<string, object?>
   {
      ["outputKey"] = "outputValue"
   };
});
