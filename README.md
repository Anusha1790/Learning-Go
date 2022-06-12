# Learning-Go

### Highlights of current commit

- Passing dynamic data from handlers to templates or page

- To solve <span style="color:red">import cycle</span> issue:

  - make models package

- <u><em>models</em></u> pkg will store models like template data, database models.

- always pass data types/structs using <u><b>pointers</b></u> (in most cases)
