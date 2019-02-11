# Yandex UI School; Lesson 1

    Chrome Dev Tools for adaptive purposes 4LIFE

### Problems with smartphones:
 * Headers don't fit
 * Tiny font
 * Tiny logo
 * Tiny buttons, can't reach

#### Problem #1 

`Things are tiny`

**Easy desision** = *viewport*

    Viewport = viewable part of a page
```html 
<meta name="viewport" content="width=device-width" />
```

Fixing this causes UI to become much bigger, as typically browser expects us to have a monitor ~900 pxls width. So this is kinda fixated zoom

###### Some viewport magic
* width 
* initial-scale=1 
* minimum-scale/maximum-scale = 1
* user-scalable = no

#### Problem #2
`Can't see everything, having horizontal scroll`

**Easy desision** = Never adjust fixated width, em helps too

BAD | SWEET
------------ | -------------
```.section_Content {width: 1200px;} ```| ```.section_Content {max-width: 1200px;} ```

#### Problem 3 
`Too big headers for phones`

**Easy desision** = CSS Media Queries

```css
 @media and (max-width:400px) {
    .header_title {
        font-size: 36px;
    }
 }
 ```
    
    ``` @media > media type > if {}```

**Breakpoints where content breaks**
* Make screen smaller
* Define when bad things start to happen
* Media those *biscuits*

#### Problem 4
`Words moving out lines`

**Easy desision** = min-width

```css
.message {
    width: 33%;
    min-width: 265px; !!
}
```

**Better Alternative** = @media
```css
.message {
    min-width: 265px;
}
@media (max-width:550px) {
    .message {
        width: 100%;
    }
}
```
Well, in horizontal that sucks

**Even Better Alternative** = *Flexbox*
```css
.header_messages {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
}

.message {
    max-width: 360px;
    flex-grow: 1; /* you can grow, brother */
    flex-shrink: 1; /* you can shrink, brother */
    flex-basis: 260px; /* Well, I want you to be this
    /* flex: 1 1 260px; */
}
```

#### Flexbox
* Adaptive lists
* Centering
* ungridable grids

*Check out:[Flexbox Cheatsheet](http://jonibologna.com/flexbox-cheatsheet/) & [Visual Guide](https://scotch.io/tutorials/a-visual-guide-to-css3-flexbox-properties)*


# Much better solution: CSS Grid Layout

    Flexbox controls just one axis, wheres CSS Grid can manage both, hence you can

### Vocabulary
* Track
* Row
* Column
* Grid line
* Grid area [rectabgle]
* Gutter [Gaps]

### Debugging

Chrome Dev Tools // Firefox ?

### Naming

```html
<div class='layout'>
    <header>HEADER</header>
    <main>CONTENT</main>
    <section>
        <header>LINKS</header>
        <h3>TAGS</h3>
        <h3>CATEGORIES</h3>
        <h3>SOCIALS</h3>
    </section>
    <footer>FOOTER</footer>
</div> 
```
We get:

![This](https://i.ibb.co/C8NH0HR/Screenshot-2019-02-11-at-22-37-10.png)

CSS Can do stuff
```css
.layout_explicit {
    display: grid; /* imma manager */
    grid-template-columns: repeat(12, 1fr); /* 12 cols in size 1 fraction unit (1/12 of )container) */
    grid-template-rows: 128px 1fr 120px; /* 3 rows with these sizes */
}
.layout_explicit > header {
    grid-column: 1 / 13; /* From 1 to end */
    /* grid-column: span 12; All 12*/ 
}
.layout_explicit > main {
    grid-column: 1 / 10;
}
.layout_explicit > section {
    grid-column: 10/13;
}
.layout_explicit > footer {
    grid-column: 1/13;
}
```
We get:

![This](https://i.ibb.co/GCfnhN8/Screenshot-2019-02-11-at-22-38-00.png)

```css
.layout_explicit-area {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    grid-template-rows: 128px 1fr 120px; 
    grid-template-areas: "h h h h"
                         "c c c s"
                         "f f f f"; /* you get it?*/
}
.layout_explicit-area > header {
    grid-area: h;
}
.layout_explicit-area > main {
    grid-area: c;
}
.layout_explicit-area > section {
    grid-area: s;
}
.layout_explicit-area > footer {
    grid-area: f;
}
```
We get:

![This](https://i.ibb.co/vjgwP07/Screenshot-2019-02-11-at-22-57-06.png)

### Case 2: Uncertain

    When we don't know how much elements we do have

```css
.layout_inplicit {
    display: grid;
    grid-template-columns: repeat(4,50px);
    grid-auto-rows: 50px; /*every 50 */
    grid-gap: 8px; /* between cells */
}
.item {
    grid-row: span 1;
}
.item_size_s {
    grid-column: span 1;
}
.item_size_m {
    grid-column: span 2;
}
.item_size_l {
    grid-column: span 3;
}
```

```css
.lay { 
    grid-template-columns: repeat(auto-fill, 50px); /* changes amount of cols when stretch  MAAGIC*/

    grid-auto-flow: dense; /* EVEN MORE MAGIC */
}
```
We get:

![This](https://i.ibb.co/g6DPsYR/Screenshot-2019-02-11-at-23-33-49.png)


#### CSS Grid
* Aligning containers
  * align-content, justify-content, justify-items, align-items, justify-self, align-self
* Auto-sizing
  * repeat(), minmax(), auto-fill, auto-fit, max-content, min-content, fit-content()
* Naming grid lines

 **Read**: *[CSS Grid Mozilla Docs](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Grid_Layout), [Interactive Course](https://scrimba.com/g/gR8PTE), [Examples](https://gridbyexample.com/), [Standards](https://drafts.csswg.org/css-grid), [Lesson Examples](https://codepen.io/dndushkin/pen/qxGWKw)*

## Extra

### Adaptive Images
Changing img size cuz

* PPI
* Width optimization

#### Some new img magic

```html
<img srcset="secretphoto.jpg 320w, 
             bigsecretphoto.jpg 800w" 
      sizes="(max-width: 320px) 280px, 800px"
      src="bigsecretphoto.jpg">

<picture>
  <source type="image/webp" srcset="pyramid.webp">
  <img src="pyramid.png">
</picture>
```

### Adaptive Typography
Two approaches:

* @media + rem
* calc() + vw

One | Two
------------ | -------------
```css
html {
  font-size: 14px;
}
@media (min-width: 768px) {
  html { font-size: 12px }
}
h1 { font-size: 2rem }
```
|
```css
h1 { font-size: 22px }

@media (min-width: 768px) {
  h1 {font-size: calc(4.4vw - 10px) }
}
```

## Courses

* **[Udacity — Mobile Web Development](https://www.udacity.com/course/mobile-web-development--cs256)**
* **[Google — Responsive Web Design Fundamentals](https://www.udacity.com/course/responsive-web-design-fundamentals--ud893)**
* **[Пацев — CSS](https://events.yandex.ru/lib/talks/1523/)**
* **[Пацев — Text](https://events.yandex.ru/lib/talks/1524/)**
* **[Пацев — Visual Formatting](https://events.yandex.ru/lib/talks/1548/)**
* **[Пацев — Layouts. At-rules](https://events.yandex.ru/lib/talks/1556/)**
* **[Пацев — Colurs & Back. Transition. Animation](https://events.yandex.ru/lib/talks/1557/)**
* **[Intro to viewport and new RWD CSS](https://vimeo.com/93347108)**