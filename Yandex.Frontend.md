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

# Problem #2
`Can't see everything, having horizontal scroll`

Never adjust fixated width

BAD | SWEET
------------ | -------------
```css .section_Content {width: 1200px;} ```| ```.section_Content {max-width: 1200px;} ```

