package fuzz_surf_jsonutil

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/isan-rivkin/surf/lib/common/jsonutil"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                jsonutil.NewJsonContainerFromBytes(data)
                return 0

            case 1:
                m := make(map[string]string)
                fuzzConsumer.FuzzMap(&m)

                rootKey, _ := fuzzConsumer.GetString()

                jsonutil.NewJsonContainerFromMap(rootKey, m)
                return 0

            case 2:
                m := make(map[string]any)
                fuzzConsumer.FuzzMap(&m)

                rootKey, _ := fuzzConsumer.GetString()

                jsonutil.NewJsonContainerFromInterface(rootKey, m)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}