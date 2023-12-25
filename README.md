# goStructProject

## json.Marshal(data)
- it will only extract and convert struct data which is publicly available
- only fields inside a struct which are publicly available will be present in the json

## struct tags
- struct tags will help us in json Marshal and unmarshal by giving the name we want for the field which should be stored in the field.