# A Visual Guide To CSS3 Flexbox Properties

[This](https://scotch.io/tutorials/a-visual-guide-to-css3-flexbox-properties) guide.

**Flexbox** –> smaller components of project
**CSS Grids** –> a whole layout

1. [Basics](#1)
2. [Usage](#2)
3. [FB container props](#3)
    1. [flex-direction](#4)
    2. [flex-wrap](#5)
    3. [flex-flow](#6)
    4. [justify-content](#7)
    5. [align-items](#8)
    6. [align-content](#9)
    7. [Note for flex containers](#10)
4. [FB items props](#11)
    1. [order](#12)
    2. [flex-grow](#13)
    3. [flex-shrink](#14)
    4. [flex-basis](#15)
    5. [flex](#16)
    6. [align-self](#17)
    7. [Note for flex items](#18)

[Check this playground](https://codepen.io/justd/pen/yydezN/)!

<a name="1"></a>

## Basics

There are *flex container* == **parent**, and *flex items* == **children**

![](https://scotch-res.cloudinary.com/image/upload/dpr_1,w_650,q_auto:good,f_auto/media/https://cask.scotch.io/2015/04/CSS3-Flexbox-Model.jpg)

It's first draft were made in 2009 and is under development.

### Browser support

For 2019 it's pretty sweet. IE11 has bugs, otherwise it's perfect.

<a name="2"></a>

## Usage

FB is set with `display` property in container. All children automatically become flex items:

```css
.flex-container {
  display: -webkit-flex; /_ Safari _/
  display: flex;
}
```

We can split properties in the two sections: *container & items properties*

<a name='3'></a>

## FB container props

<a name='4'></a>

### flex-direction

Sets the direction of main axis: elements can be laid in rows or in columns:

```css
.flex-container {
  /* In row from left to right */
  -webkit-flex-direction: row; /_ Safari _/
  flex-direction:         row;

  /* In row from right to left */
  flex-direction:         row-reverse;

  /* Top to bottom and reversed */
  flex-direction:         column;
  flex-direction:         column-reverse;
}
```

**Default**: `row`

`row` & `row-reverse` depend on writing mode.

<a name='5'></a>

### flex-wrap

Sets the amount of lines. Should it be a single or multiple lines, and direction.

```css
.flex-container {
  /* In one row */
  -webkit-flex-wrap: nowrap; /_ Safari _/
  flex-wrap:         nowrap;

  /* In multiple rows */
  flex-wrap:         wrap;
  flex-wrap:         wrap-reverse;
}
```

**Default**: `nowrap`

Same with writing mode.

<a name='6'></a>

### flex-flow

Shorthand for `flex-direction` and `flex-wrap`.

```css
.flex-container {
  -webkit-flex-flow:  flex-direction flex-wrap ; /_ Safari _/

  /* For example */
  flex-flow:          row wrap;
}
```

**Default**: `row nowrap`

<a name='7'></a>

### justify-content

Aligns items alongs the main axis, distributes space.

```css
.flex-container {
  /* Aligned to the left side of the container */
  -webkit-justify-content: flex-start; /_ Safari _/
  justify-content:         flex-start;

  /* Aligned to the right */
  justify-content:         flex-end;

  /* In the center of the container */
  justify-content:         center;

  /* First and last items are at edges, others fill the space with equal gap between them */
  justify-content:         space-between;

  /* Same, but also gap before first item and after last */
  justify-content:         space-around;
}
```

**Default**: `flex-start`

<a name='8'></a>

### align-items

Same as `justify-content`, but in perpendicular direction.

```css
.flex-container {
  /* Items fill the whole axis length from cross start to cross end */
  -webkit-align-items: stretch; /_ Safari _/
  align-items:         stretch;

  /* Items are stacked to the cross start or end */
  align-items:         flex-start;
  align-items:         flex-end;

  /* Items are stacked in the center */
  align-items:          center;

  /* By items' baselines */
  align-items:          baseline;
}
```

**Default**: `stretch`

<a name='9'></a>

### align-content

On multiple lines sets the gaps between lines:

```css
.flex-container {
  /* Items are displayed with distributed space after every row of flex items */
  -webkit-align-content: stretch; /_ Safari _/
  align-content:         stretch;

  /* Items are stacked toward the cross start  or end */
  align-content:         flex-start;
  align-content:         flex-end;

  /* Same to justify */
  align-content:         space-between;
  align-content:         space-around;
}
```

**Default**: `stretch`

<a name='10'></a>

### Note for flex containers

* all of the `column-*` properties have no effect on a flex container.
* the `::first-line` and `::first-letter` pseudo-elements do not apply to flex containers.

<a name="11"></a>

## FB items props

<a name="12"></a>

### order

Order in which children are displayed.

```css
.flex-item {
  -webkit-order: ; /_ Safari _/
  order:         ;
}
```

**Default**: 0

<a name="13"></a>

### flex-grow

Decides, wether this item can grow

```css
.flex-item {
  -webkit-flex-grow: ; /_ Safari _/
  flex-grow:         ;
}
```

**Default**: 0

<a name="14"></a>

### flex-shrink

Decides, wether this item can shrink

```css
.flex-item {
  -webkit-flex-shrink: ; /_ Safari _/
  flex-shrink:         ;
}
```

**Default**: 1

<a name="15"></a>

### flex-basis

Initial main size of the item, same as width while `flex-direction: row`, same as height, while `column`:

```css
.flex-item {
  -webkit-flex-basis: auto ; /_ Safari _/
  flex-basis:         auto ;
}
```

**Default**: `auto`

<a name="16"></a>

### flex

Shorthand for `flex-grow`,`flex-shrink` and `flex-basis`:

```css
.flex-item {
  -webkit-flex:1 1 200px; /_ Safari _/
  flex:        1 1 200px;
}
```

**Default**: `0 1 auto`

<a name='17'></a>

### align-self

Overrides `align-items` for this item.

```css
.flex-item {
  -webkit-align-self: auto | flex-start | flex-end | center | baseline | stretch; /_ Safari _/
  align-self:         auto | flex-start | flex-end | center | baseline | stretch;
}
```

**Default**: `auto`

<a name='18'></a>

### Note for flex items

* `float`, `clear` and `vertical-align` have no effect on a flex item, and do not take it out-of-flow.