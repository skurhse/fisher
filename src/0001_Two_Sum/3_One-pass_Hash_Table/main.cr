module OnePass
  def self.two_sum(nums, target)
    compliment_indices = {} of Int32 => Int32

    nums.each_with_index do |num, index|
      if compliment_indices.has_key?(num)
        return compliment_indices[num], index
      end

      compliment_indices[target - num] = index
    end
  end
end
