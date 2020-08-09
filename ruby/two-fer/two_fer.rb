=begin
Write your code for the 'Two Fer' exercise in this file. Make the tests in
`two_fer_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/two-fer` directory.
=end

module TwoFer 
    MESSAGE = "One for %s, one for me."
    def self.two_fer(message="you")
        MESSAGE % message
    end
end
