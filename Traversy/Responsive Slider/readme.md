# Responsive Image Slider No libs

## I. HTML

1. First of all, really simple html:

```html
<div class="slider">
  <div class="current slide"></div>
  <div class="slide"></div>
  <div class="slide"></div>
</div>
```

2. Then add buttons after that:

```html
<div>
  <button id="prev"><</button>
  <button id="next">></button>
</div>
```

3. Lastly, connect js file:

```html
<script src="slider.js"></script>
```

That's it. Now let's do css.

## II. CSS

1. Let's make our slider relative

```css
.slider {
  position: relative;
  overflow: hidden;
  width: 100vw;
}
```

2. So slides can be absolute. Opacity to 1 just for the current slide

```css
.slide {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  opacity: 0;
  transition: opacity 0.5s ease-in-out;
}

.slide.current {
  opacity: 1;
}
```

3. Then give images via bg:

```css
.slide:first-child {
  background: url(assets/1.jpg) no-repeat center center/cover;
}

.slide:nth-child(2) {
  background: url(assets/2.jpg) no-repeat center center/cover;
}

.slide:nth-child(3) {
  background: url(assets/3.jpg) no-repeat center center/cover;
}
```

## III. JS

1. Set the vars:

```js
const slides = document.querySelectorAll(".slide");
const next = document.querySelector("#next");
const prev = document.querySelector("#prev");

const auto = false;
const intervalTime = 5000;
let slideInterval;
```

2. then onClick funcs:

```js
const nextSlide = () => {
  const current = document.querySelector(".current");
  current.classList.remove("current");

  if (current.nextElementSibling)
    current.nextElementSibling.classList.add("current");
  else slides[0].classList.add("current");

  setTimeout(() => current.classList.remove("current"));
};
```

```js
const prevSlide = () => {
  const current = document.querySelector(".current");
  current.classList.remove("current");

  if (current.prevElementSibling)
    current.prevElementSibling.classList.add("current");
  else slides[slides.length - 1].classList.add("current");

  setTimeout(() => current.classList.remove("current"));
};
```

3. Then events:

```js
next.addEventListener("click", event => nextSlide());
prev.addEventListener("click", event => prevSlide());
```

4. In order to automate just do:

```js
if (auto) sliderInterval = setInterval(nextSlide, intervalTime);
```

5. Yet we need to reset the interval if the button is clicked. we have to modify `click` events

```js
next.addEventListener("click", event => {
  nextSlide();
  if (auto) {
    clearInterval(slideInterval);
    sliderInterval = setInterval(nextSlide, intervalTime);
  }
});
prev.addEventListener("click", event => {
  prevSlide();
  if (auto) {
    clearInterval(slideInterval);
    sliderInterval = setInterval(nextSlide, intervalTime);
  }
});
```

## [Result](https://codesandbox.io/s/lp6w1mxjxm)
