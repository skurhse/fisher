#!/usr/bin/env crystal

# REQ: Defines the TreeNode class for tree problems. <skr 2022-01-04>

class TreeNode
  property value : Int32?
  property left : TreeNode?
  property right : TreeNode?

  def initialize(values : Array(Int32?))
    return if values.size == 0
    @value = values[0]
    return if values.size == 1

    left_values = [values[1]]
    right_values = [values[2]]
    values[3..].each_with_index do |value, index|
      sub_index = index % 4

      if [0,1].includes?(sub_index)
        left_values.push(value)
      else
        right_values.push(value)
      end
    end

    @left = TreeNode.new(left_values)
    @right = TreeNode.new(right_values)
  end

  def self.[](*values)
    self.new(values.to_a)
  end

  def self.[]
    self.new([] of Int32)
  end
end
