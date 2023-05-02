using BenchmarkDotNet.Running;

class Program
{
    static void Main(string[] args) 
        => BenchmarkRunner.Run<Xor.Benchmarks.DecodeBenchmark>();
}
