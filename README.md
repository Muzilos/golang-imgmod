# golang-imgmod

## Usage
```bash
go build main.go
go run ./main [--input <input_file_name>] [--output <output_file_name>] [--type <jpeg/bmp>] [--strength <float64>]
```

### Arguments
``` bash
  --input string
      image file to read (default "./img/painting.jpg")
  --output -o string
      filename for output (default "./img/modded.jpg")
  --type -t string
      file format (jpeg, bmp) (default "jpeg")
  --strength -s float64
      strength of modification (default 3.0)
```

## Contact
Please don't contact me