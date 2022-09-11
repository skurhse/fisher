require "spec"
require "../main.cr"

describe OnePass do
  describe "#two_sum" do
    it "returns the indices of two numbers with sum target" do
      OnePass.two_sum([2, 7, 11, 15], 9).should eq ({0, 1})
      OnePass.two_sum([3, 2, 4], 6).should eq ({1, 2})
      OnePass.two_sum([3, 3], 6).should eq ({0, 1})
    end
  end
end
