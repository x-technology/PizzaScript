number = Random.rand(10)
puts "guess what I've guessed!"

guess = gets.chomp.to_i
while guess != number
  puts "nah, wrong. try one more time"
  guess = gets.chomp.to_i
end

puts "just lucky"
