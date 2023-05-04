package fuzz_surf_proto

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/isan-rivkin/surf/lib/common/proto"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                var acc proto.Accomulator
                fuzzConsumer.GenerateStruct(&acc)

                delim, _ := fuzzConsumer.GetString()

                acc.ToString(delim)
                return 0

            case 1:
                var acc proto.Accomulator
                fuzzConsumer.GenerateStruct(&acc)

                b, _ := fuzzConsumer.GetBytes()

                proto.ParseUnknown(b, &acc)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}