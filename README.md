# Assignment

Create a rails app the uses a provided binary to generate a graphical website hit counter.

The binary is compiled for Linux, Windows, and Darwin (MacOS) in 32bit (386) and 64bit (amd64) architectures.  The binary can be downloaded here:

[Image server binaries](#)

### Example binary usage:

```
# in one terminal
./imgsrc-darwin-amd64 -port=8088
```

```
# in a second terminal
curl localhost:8088 -d '{"canvas_width": 200, "canvas_height": 200, "img_commands": [
{"cmd": "SetRGBA", "args": [1,1,0,1]},
{"cmd": "SetLineWidth", "args": [4]},
{"cmd": "DrawLine", "args":[20, 20, 40, 40]},
{"cmd": "SetRGBA", "args": [1, 0, 1, 0.8]},
{"cmd": "DrawCircle", "args": [30, 30, 30]},
{"cmd": "Fill"}, {"cmd": "SetRGBA", "args": [0,0,1,0.6]},
{"cmd": "DrawCircle", "args": [30, 30, 30]},
{"cmd": "Stroke"}
]}' > img.png
```

This produces a PNG formatted image.  As you can see the height and width of the canvas can be specified along with a list of drawing commands.

### Supported Drawing Commands

- `SetRGBA(red float 0-1, green float 0-1, blue float 0-1, alpha float 0-1)`

Sets the color using the supplied RGB and Alpha values.

- `DrawCircle(x float, y, float, radius float)`

- `SetLineWidth(width float)`

Sets the line width for `DrawLine` and `Stroke` commands

- `DrawRectange(x1 float, y1 float, x2 float, x2 float)`

Creates a rectangle path on the canvas.  This must be followed by a `Fill` or `Stroke` to fill and draw an outline of the rectangle.

- `Fill()`

Fills the circle or rectangle path created by `DrawCircle` or `DrawRectange`

- `Stroke()`

Outlines the circle or rectangle path created by `DrawCircle` or `DrawRectange`

- `DrawLine(x1 float, y1 float, x2 float, x2 float)`

Draws a line from x1,y1 to x2,y2.  There is no need to follow this command with `Fill` or `Stroke` commands.

# Solution

A solution should be a rails app that has a page or pages with an image counter.  The counter should increase for each request.  The rails app does not need to start or stop the supplied image server binary.  This can be done manually.  Provide us with a zipped up copy of your rails app.

An example web page might look like this:

![example](example.png)
