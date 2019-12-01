defmodule Santa do
  def calculate_fuel_cost(mass) do
    case div(mass, 3) - 2 do
      fuel when fuel <= 0 ->
        0

      fuel ->
        fuel + calculate_fuel_cost(fuel)
    end
  end
end

IO.read(:stdio, :all)
|> String.split("\n", trim: true)
|> Enum.map(&String.to_integer/1)
|> Enum.map(&Santa.calculate_fuel_cost/1)
|> Enum.sum()
|> IO.inspect(label: "Total fuel required")
