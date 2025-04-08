# imgembed
This Github Action takes an HTML file and replaces all `<img>` tags referencing a file and replaces the `src` attribute with the base64 encoded representation of the image files content. Image files are searched in the working directory. If a image file cannot be found in the filesystem the action terminates with an error.

## Usage in workflow:

```
- uses: faryon93/imgembed@v1
  with:
    input_file: "in.html"          # filepath to input file
    output_file: "out.html"        # filepath to output file
```
