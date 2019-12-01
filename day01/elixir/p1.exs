IO.read(:stdio, :all)
|> String.split("\n", trim: true)
|> Enum.map(&String.to_integer/1)
|> Enum.map(fn x -> div(x, 3) - 2 end)
|> Enum.sum()
|> IO.inspect(label: "Total fuel required")
