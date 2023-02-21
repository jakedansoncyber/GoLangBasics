using System.Diagnostics;
using System.Threading;
// See https://aka.ms/new-console-template for more information
Stopwatch stopwatch = new Stopwatch();
stopwatch.Start();
var num = new Random();

var list = new List<int>();
for(int i = 0; i < 100000; i++){
    list.Add(num.Next(2147452437));
}

list.Sort();

for (int i = 0; i < 100000; i++){
    Console.WriteLine(list[i]);
}
System.Console.WriteLine($"Done in {stopwatch.ElapsedMilliseconds}");
stopwatch.Stop();


