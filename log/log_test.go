package log_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	ctxutil "github.com/tyeryan/l-protocol/context"
	. "github.com/tyeryan/l-protocol/log"
	"testing"
)

func TestLog(t *testing.T) {
	log := GetLogger("TestLog")
	ctx := ctxutil.NewContext(ctxutil.WithStan("ThisIsStan"))
	ctx = ctxutil.Add(ctx, ctxutil.YBPCurrentOperator, "aayush@you.co")

	name := "s1name"
	log.Add("userID", "123456")
	log.Infow(ctx, "this is a log")
	log.Infow(ctx, "this is a log with things", "key1", 1234)
	log.Infow(ctx, "this is a log with things with dot(.) in key", "key.KAY", 1234)
	log.Infow(ctx, "this is a log with struct", "key1", 1234, "keyStruct", s1{name: &name, value: 9999, values: []string{"v1", "v2", "v3"}})

	log.Add("youID", "Y-12278623")
	log.Errorw(ctx, "this is a error")
	log.Errore(ctx, "this is a error with things", fmt.Errorf("errrrrrrrror"), "key1", 1234)
	log.Errore(ctx, "this is a error with struct", fmt.Errorf("errrrrrrrror"), "key1", 1234, "keyStruct", s1{name: &name, value: 9999, values: []string{"v1", "v2", "v3"}})

	log.Alertw(ctx, "this is a alert")
	log.Alerte(ctx, "this is a alert with things", fmt.Errorf("errrrrrrrror"), "key1", 1234)
	log.Alerte(ctx, "this is a alert with struct", fmt.Errorf("errrrrrrrror"), "key1", 1234, "keyStruct", s1{name: &name, value: 9999, values: []string{"v1", "v2", "v3"}})
}

func TestCanonicalLogWithErrorHandling(t *testing.T) {

	testErrorFunc := func(inErr error) (err error) {
		log := GetLogger("TestCanonicalLog")
		ctx := ctxutil.NewContext(ctxutil.WithStan("TestCanonicalLogWithErrorHandling"))
		log.Add(RequestId, "1234")

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in func", r)
			}
			log.Canonical(ctx, "this is a Canonical log", err, recover())
		}()

		// test to see how Canonical display for struct / array
		type Card struct {
			CardID    string
			AccountNo string
			YouID     string
		}
		c := Card{CardID: "cardId", AccountNo: "12786237862346784", YouID: "Y-8798897"}
		cards := make([]Card, 0)
		cards = append(cards, c)
		log.Add("c", c, "cards", cards)

		return inErr
	}

	assert.NotNil(t, testErrorFunc(errors.New("test error")))
	assert.Nil(t, testErrorFunc(nil))
}

func TestCanonicalLogCustomPanicHandleFurtherPanic(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()

	testPanicFunc := func() *Log {
		log := GetLogger("TestCanonicalLogCustomPanicHandle")
		ctx := ctxutil.NewContext(ctxutil.WithStan("TestCanonicalLogCustomPanicHandle"))
		log.Add(RequestId, "1234")

		defer func() {
			log.Canonical(ctx, "this is a Canonical log", nil, recover())
		}()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in func", r)
				panic(r)
			}
		}()

		panic("test panic")

		return log
	}
	testPanicFunc()
}

func TestCanonicalLogCustomPanicHandleNoFurtherPanic(t *testing.T) {
	defer func() {
		assert.Nil(t, recover())
	}()

	testPanicFunc := func() {
		log := GetLogger("TestCanonicalLogCustomPanicHandleNoFurtherPanic")
		ctx := ctxutil.NewContext(ctxutil.WithStan("TestCanonicalLogCustomPanicHandleNoFurtherPanic"))
		log.Add(RequestId, "1234")

		defer func() {
			log.Canonical(ctx, "this is a Canonical log", nil, recover())
		}()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in func", r)
			}
		}()

		panic("test panic")

	}

	testPanicFunc()
}

func TestCanonicalLogDefaultUsage(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()

	testPanicFunc := func(doPanic bool) {
		log := GetLogger("TestCanonicalLogDefaultUsage")
		ctx := ctxutil.NewContext(ctxutil.WithStan("TestCanonicalLogDefaultUsage"))
		log.Add(RequestId, "1234")

		defer func() {
			// recover pass to log, will mark YouPanic and propagate panic out
			log.Canonical(ctx, "this is a Canonical log", nil, recover())
		}()

		if doPanic {
			panic("test panic")
		}
	}

	testPanicFunc(false)
	testPanicFunc(true)
}

type s1 struct {
	name   *string
	value  int64
	values []string
}
