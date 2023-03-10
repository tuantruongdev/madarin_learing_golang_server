
# Madarin Learning Golang Sever

An small & simple project for chinese learner 

## Note
This project is backend side for[mandarin_learning_android_client](https://github.com/tuantruongdev/mandarin_learning_android_client)
![Deploy diagram](https://raw.githubusercontent.com/tuantruongdev/mandarin_learning_android_client/addab76d1b998fc3e11b10b8c67e7754ac713db9/app/src/main/res/drawable/Deploy%20diagram.png)
## Authors

- [@tuantruongdev](https://github.com/tuantruongdev)

## API Reference

#### Word lookup

```http
  GET /api/v1/word/lookup/:character
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `character`      | `string` | **Required** Chinese character |


####  Word Example lookup

```http
  GET /api/v1/word/sentences/:character?&level=Newbie&includeAudio=1
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `character` | `string` | **Required** Chinese character  |
| `level` | `string` |Newbie,Elementary,Pre Intermediate,Intermediate,Upper-Intermediate |
| `includeAudio` | `number` | `1` for getting audio, `0` for none |


## Related

Here are some related projects

Big thank to [Pinyin-word-api](https://github.com/felipemarinho97/pinyin-word-api)


