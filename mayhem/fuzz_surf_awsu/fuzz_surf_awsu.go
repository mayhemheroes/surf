package fuzz_surf_awsu

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/isan-rivkin/surf/lib/awsu"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                var InputArr []*awsu.AWSSessionInput
                repeat, _ := fuzzConsumer.GetInt()

                for i := 0; i < repeat; i++ {

                    var temp awsu.AWSSessionInput
                    fuzzConsumer.GenerateStruct(&temp)

                    InputArr = append(InputArr, &temp)
                }

                awsu.NewSessionInputMatrix(InputArr)
                return 0

            case 1:
                profile, _ := fuzzConsumer.GetString()
                region, _ := fuzzConsumer.GetString()

                awsu.NewSessionInput(profile, region)
                return 0

            case 2:
                var Input awsu.AuthInput
                fuzzConsumer.GenerateStruct(&Input)

                awsu.NewACM(&Input)
                return 0

            // case 3:
            //     var Input awsu.AuthInput
            //     fuzzConsumer.GenerateStruct(&Input)

            //     awsu.NewS3(&Input)
            //     return 0

            // case 4:
            //     var Input awsu.AuthInput
            //     fuzzConsumer.GenerateStruct(&Input)

            //     awsu.NewDDB(&Input)
            //     return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}