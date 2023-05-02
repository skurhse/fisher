using System;
using BenchmarkDotNet.Attributes;

namespace Xor.Benchmarks;

public class DecodeBenchmark
{
    private int[] encoded;

    [Params(1000, 10000)]
    public int N;

    [GlobalSetup]
    public void Setup()
    {
        var random = new Random();
        encoded = new int[N];
        for (int i = 0; i < N; i++)
        {
            encoded[i] = random.Next();
        }
    }

    [Benchmark]
    public int[] Decode() => Xor.Decode(encoded, 1);

    [Benchmark]
    public int[] DecodeWithSpans() => Xor.DecodeWithSpans(encoded, 1);

    [Benchmark]
    public int[] DecodeWithSpansWithoutHoist() => Xor.DecodeWithSpansWithoutHoist(encoded, 1);
}
