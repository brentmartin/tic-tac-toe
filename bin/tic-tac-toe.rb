@board = [1, 2, 3, 4, 5, 6, 7, 8, 9]
@player1 = []
@player2 = []
@p1_win = FALSE
@p2_win = FALSE


def p1_turn
  def move(n)
    if @board.include?(n)
      @player1.push(n)
      @board.delete(n)
    else
      puts "Play not available, please try again!"
      p1_turn
    end
  end
  puts "Player 1, what have you?"
  choice = gets.chomp.to_i
  move(choice)
end

def p2_turn
  def move(n)
    if @board.include?(n)
      @player2.push(n)
      @board.delete(n)
    else
      puts "Play not available, please try again!"
      p2_turn
    end
  end
  puts "Player 2, what have you?"
  choice = gets.chomp.to_i
  move(choice)
end

def p1_finnish(a, b, c)
  if @player1.include?(a) && @player1.include?(b) && @player1.include?(c)
    @p1_win = true
    puts
    puts "Player 1 wins!"
    puts
    puts "Board now has available: #{@board}"
    puts "Player 1 now has: #{@player1}"
    puts "Player 2 now has: #{@player2}"
    puts
    p2_turn
  end
end

def p2_finnish(a, b, c)
  if @player2.include?(a) && @player2.include?(b) && @player2.include?(c)
    @p2_win = true
    puts
    puts "Player 2 wins!"
    puts
    puts "Board now has available: #{@board}"
    puts "Player 1 now has: #{@player1}"
    puts "Player 2 now has: #{@player2}"
    puts
    p1_turn
  end
end

def p1_win_check
  unless @p1_win || @p2_win
    puts "Firing up the old determiner"
    p1_finnish(1,2,3)
    p1_finnish(4,5,6)
    p1_finnish(7,8,9)
    p1_finnish(1,4,7)
    p1_finnish(2,5,8)
    p1_finnish(3,6,9)
    p1_finnish(1,5,9)
    p1_finnish(3,5,7)
  end
end

def p2_win_check
  unless @p1_win || @p2_win
    puts "Firing up the old determiner"
    p2_finnish(1,2,3)
    p2_finnish(4,5,6)
    p2_finnish(7,8,9)
    p2_finnish(1,4,7)
    p2_finnish(2,5,8)
    p2_finnish(3,6,9)
    p2_finnish(1,5,9)
    p2_finnish(3,5,7)
  end
end

puts "LETS GET THIS PARTY STARTED"
puts
puts "Board starts with: #{@board}"
puts "Player 1 starts with: #{@player1}"
puts "Player 2 starts with: #{@player2}"
puts


until @board == []
  p1_turn
  p1_win_check
  system("clear")

  puts
  puts "End of player 1's turn"
  puts "Board now has available: #{@board}"
  puts "Player 1 now has: #{@player1}"
  puts "Player 2 now has: #{@player2}"
  puts

  p2_turn
  p2_win_check
  system("clear")

  puts
  puts "End of player 2's turn"
  puts "Board now has available: #{@board}"
  puts "Player 1 now has: #{@player1}"
  puts "Player 2 now has: #{@player2}"
  puts

end
