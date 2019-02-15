
[< Previous abstract](React__8.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React) | [Next abstract >](React__10.md)
----------------------- | ----------------------------|-----------------------------

## 9. Forms

  HTML form elements work a little bit differently from other DOM elements in React, because form elements naturally keep some internal state. For example, this form in plain HTML accepts a single name:

```html
<form>
  <label>
    Name:
    <input type="text" name="name" />
  </label>
  <input type="submit" value="submit" />
</form>
```

It works in HTML in case you want user to submit and go to another page. If you want this in React — it works. However, in most cases JS functions are more convinient. Best way to do such in React called "controlled component".