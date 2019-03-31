# Smooth Scroll

## Pure CSS

```css
scroll-behavior: smooth;
```

Also snapping to the viewports!

But no Safari, IE, Edge :c

1. Remove scrollbars from `body`:

```css
body {
  overflow: hidden;
}
```

2. However, container/wrapper should be able to be scrolled:

```css
.container {
  overflow-y: scroll;
  scroll-behavior: smooth;
}
```

3. Snap effect!

```css
.container {
  scroll-snap-type: y mandatory;
}
```

And in children:

```css
.container > article {
  scroll-snap-align: center;
}
```

**I love this so much!**

## II jQuery (that depricated thing, you know)

1. connect **minified**, not slim version of jQuery, then do:

```js
$(".navbar a").on("click", function(event) {
  if (this.hash !== "") {
    //this.hash stores for ex "#about"
    event.preventDefault();

    const hash = this.hash;

    $("html", "body").animate(
      {
        scrollTop: $(hash).offset().top
      },
      800
    );
  }
});
```

## III but what about Safari?

Don't tell anybody, there is a [lib](https://github.com/cferdinandi/smooth-scroll)!..

1. Connect via CDN:

```html
<script src="https://cdn.jsdelivr.net/gh/cferdinandi/smooth-scroll@15.0.0/dist/smooth-scroll.polyfills.min.js" />
```

2. Initialise smooth scroll:

```js
const scroll = new SmoothScroll('a[href*="#"]', {
  speed: 800
});
```

That's it. No snaps!
