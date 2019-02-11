# Yandex UI School

## Lesson 1

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

##### Flexbox
* Adaptive lists
* Centering
* ungridable grids

*Check out:[Flexbox Cheatsheet](http://jonibologna.com/flexbox-cheatsheet/) & [Visual Guide](https://scotch.io/tutorials/a-visual-guide-to-css3-flexbox-properties)*

## Much better solution

