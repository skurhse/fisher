namespace Xor.Tests;

public class DecodeTest
{
    [Theory]
    [InlineData(new int[3] { 1, 2, 3 }, 1, new int[4] { 1, 0, 2, 1 })]
    [InlineData(new int[4] { 6, 2, 7, 3 }, 4, new int[5] { 4, 2, 0, 7, 4 })]
    public void Examples(int[] encoded, int first, int[] expected)
    {
        Assert.Equal(Xor.Decode(encoded, first), expected);
        Assert.Equal(Xor.DecodeWithSpans(encoded, first), expected);
        Assert.Equal(Xor.DecodeWithSpansWithoutHoist(encoded, first), expected);
    }
}
