package main

import("fmt"
	   "os"
)

func main()  {

	if len(os.Args) != 3 {
		fmt.Println("error: usage input.txt output.txt")
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	data, err := os.ReadFile(file1)
	if err != nil {
		fmt.Println("Error: failed to read file")
	}
	fmt.Println(data)

	text := string(data)

	text = Capitalize(text)
	text = Lower(text)
	text = HexToDecimal(text)
	text = BinToDecimal(text)

	erro := os.WriteFile(file2, []byte(text), 0644); if err == nil {
		fmt.Println("error: file not found")
	}
	fmt.Println(erro)


}
