# Excersice 1

---

- Author: Hoplin

---

## Question

아래 조건을 만족하는 `Cafe` 메세지를 작성해보자

- Cafe
  - 좌표를 나타내는 `Axis` 타입이다. `Axis`에는 `latitude`,`longitude`값이 있어야한다.
  - 카페 이름
  - 카페 메뉴와 가격
  - 카페에 대한 해시태그들
  - 글이 작성된 요일. 단 요일은 총 7개중 하나여야한다.

## Solution

우선 각각의 필드를 분석해본다

1. 좌표를 나타내는 `Axis`는 경,위도를 의미하는 `latitude`와 `longitude` 필드를 포함하고 있다. 이에 맞춰서 `Axis` 메세지 타입을 정의해본다. 단 경,위도는 각각 부동소수점을 가지므로 `float` 혹은 `double`로 정의한다.

```proto
message Axis{
    float latitude = 1;
    float longitude = 2;
}
```

2. 카페 이름은 문자열 타입이므로 `string`타입을 가진다

```proto
string name = 2;
```

3. 카페 `메뉴`와 `가격`이 각각 매핑되기 좋은 자료형은 `map` 자료형이다. `메뉴`는 문자열 타입이고, `가격`은 정수형 타입이므로 `map<string,uint32>`타입으로 정의한다(`uint32`가 아니어도 괜찮다)

```proto
map<string,uint32> menus = 3;
```

4. 카페에 대한 해시태그들은 여러 문자열 값을 가질 수 있으므로, `repeated` 타입 성격을 지니게 된다. 추가적으로 해시태그는 각각 문자열이므로 `repeated string`으로 정의를 할 수 있다.

```proto
repeated string hashtags = 4;
```

5. 요일을 작성한다. 단 주어진 조건을 보면 7개 중 하나를 작성하라고 하였으므로 `enum`타입으로 요일을 정의해준다.

```proto
enum DayOfWeek {
    DAY_OF_WEEK_MONDAY = 0;
    DAY_OF_WEEK_TUESDAY = 1;
    DAY_OF_WEEK_WEDNESDAY = 2;
    DAY_OF_WEEK_THURSDAY = 3;
    DAY_OF_WEEK_FIRDAY = 4;
    DAY_OF_WEEK_SATURDAY = 5;
    DAY_OF_WEEK_SUNDAY = 6;
}
```

위에서 정의한 조건들을 조합하여 Protobuffer를 완성시켜본다.

```proto
syntax = "proto3";

message Axis{
    float latitude = 1;
    float longtitude = 2;
}

enum DayOfWeek {
    DAY_OF_WEEK_MONDAY = 0;
    DAY_OF_WEEK_TUESDAY = 1;
    DAY_OF_WEEK_WEDNESDAY = 2;
    DAY_OF_WEEK_THURSDAY = 3;
    DAY_OF_WEEK_FIRDAY = 4;
    DAY_OF_WEEK_SATURDAY = 5;
    DAY_OF_WEEK_SUNDAY = 6;
}

message Cafe{
    Axis axis = 1;
    string name = 2;
    map<string,uint32> menus = 3;
    DayOfWeek day = 4;
}
```
