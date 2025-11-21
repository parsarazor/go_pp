// JXA Script to create a Keynote presentation
// Run with: osascript -l JavaScript generate_keynote.js

var Keynote = Application("Keynote");
Keynote.includeStandardAdditions = true;

// Theme and Slide Setup
var themeName = "Black"; // Try to use the standard "Black" theme
var width = 1920;
var height = 1080;

// Create new presentation
var doc = Keynote.Document({
  documentTheme: Keynote.themes[themeName],
  width: width,
  height: height,
});

Keynote.documents.push(doc);

// Helper function to add a slide
function addSlide(title, body, masterName) {
  var master = doc.masterSlides[masterName || "Title & Bullets"];
  var slide = Keynote.Slide({ baseSlide: master });
  doc.slides.push(slide);

  if (title) {
    slide.defaultTitleItem.objectText = title;
  }
  if (body) {
    slide.defaultBodyItem.objectText = body;
  }
  return slide;
}

// Helper for Title Slide
function addTitleSlide(title, subtitle) {
  var master = doc.masterSlides["Title - Center"];
  var slide = Keynote.Slide({ baseSlide: master });
  doc.slides.push(slide);

  slide.defaultTitleItem.objectText = title;
  slide.defaultBodyItem.objectText = subtitle;
}

// --- Content ---

// Slide 1: Title
addTitleSlide("Standard Streams & fmt Package", "Deep Dive in Go");

// Slide 2: Agenda
addSlide(
  "Agenda",
  "1. Standard Streams (Stdin, Stdout, Stderr)\n" +
    "2. The fmt Package\n" +
    "3. Advanced Interfaces (Stringer, Formatter)\n" +
    "4. Best Practices"
);

// Slide 3: Standard Streams
addSlide(
  "1. Standard Streams",
  "The Unix Philosophy:\n\n" +
    "• Stdin (0): Standard Input (Keyboard)\n" +
    "• Stdout (1): Standard Output (Terminal)\n" +
    "• Stderr (2): Standard Error (Logs/Errors)"
);

// Slide 4: Go Implementation
addSlide(
  "Standard Streams in Go",
  "Package: os\n" +
    "Type: *os.File (Implements io.Reader & io.Writer)\n\n" +
    'os.Stdout.WriteString("Hello")\n' +
    'os.Stderr.WriteString("Error")'
);

// Slide 5: fmt Package
addSlide(
  "2. The fmt Package",
  "Three main categories:\n\n" +
    "1. Print/Scan: Stdout / Stdin\n" +
    "   (Println, Printf, Scanln)\n\n" +
    "2. Fprint/Fscan: io.Writer / io.Reader\n" +
    "   (Fprintf, Fscan)\n\n" +
    "3. Sprint/Sscan: string\n" +
    "   (Sprintf, Sscan)"
);

// Slide 6: Formatting Verbs
addSlide(
  "Formatting Verbs",
  "%v   : Default format\n" +
    "%+v  : Struct fields\n" +
    "%#v  : Go Syntax (Debug)\n" +
    "%T   : Type\n" +
    "%d   : Base 10 Integer\n" +
    "%x   : Hex\n" +
    "%s   : String\n" +
    "%p   : Pointer"
);

// Slide 7: Advanced Interfaces
addSlide(
  "3. Advanced Interfaces",
  "Control how your types are printed.\n\n" +
    "1. Stringer: For %v and %s\n" +
    "2. GoStringer: For %#v (Debug)\n" +
    "3. Formatter: Full control over all verbs"
);

// Slide 8: Best Practices
addSlide(
  "4. Best Practices",
  "• Use log for Errors: Prefer log.Println for system errors.\n" +
    "• Buffering: fmt is unbuffered. Use bufio.Writer for high volume.\n" +
    "• Input Safety: Always check errors from Scan functions."
);

// Slide 9: End
addTitleSlide("Thank You", "Q&A");

console.log("Presentation created successfully!");
