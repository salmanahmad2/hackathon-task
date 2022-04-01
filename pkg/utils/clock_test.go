package utils_test

import (
	"testing"
	"time"

	"hackathon/pkg/utils"
)

func TestNowUnix_WhenCalled_ShouldEqualUTCToUnix(t *testing.T) {
	// Arrange
	clock := utils.NewClock()

	// Act
	unix := clock.NowUnix()

	// Assert
	utc := utils.ToTime(&unix)
	expectedUnix := utils.ToUnix(utc)
	if unix != *expectedUnix {
		t.Errorf("Expected NowUnix to equal UTC transformed to Unix; got unix = %v and "+
			"transformed UTC = %v", unix, expectedUnix)
	}
}

func TestNowUTC_WhenComparedWithEST_ShouldNotBeEqual(t *testing.T) {
	// Arrange
	clock := utils.NewClock()

	// Act
	utc := clock.NowUTC()
	estLocation, _ := time.LoadLocation("EST")
	est := utc.In(estLocation)

	// Assert
	if utc == est {
		t.Errorf("Expected UTC to not equal EST; got UTC = %v and EST = %v", utc, est)
	}
}

func TestNowUTC_WhenTransformedToUnix_ShouldBeEqualToESTTransformed(t *testing.T) {
	// Arrange
	clock := utils.NewClock()
	utc := clock.NowUTC()
	estLocation, _ := time.LoadLocation("EST")
	est := utc.In(estLocation)

	// Act
	utcTimestamp := utils.ToUnix(&utc)
	estTimestamp := utils.ToUnix(&est)

	// Assert
	if utc == est {
		t.Errorf("Expected UTC to not equal EST; got UTC = %v and EST = %v", utc, est)
	}
	if *utcTimestamp != *estTimestamp {
		t.Errorf("Expected timestamps from UTC and EST to be equal; got UTC = %v and EST = "+
			"%v", utcTimestamp, estTimestamp)
	}
}

func TestToUnix_WhenValidTime_ShouldReturnValidUnixTimestamp(t *testing.T) {
	// Arrange
	validTime := time.Date(2019, time.March, 28, 18, 26, 50, 0, time.FixedZone("", 0))

	// Act
	result := utils.ToUnix(&validTime)

	// Assert
	expectedUnixTimestamp := int64(1553797610000)
	if *result != expectedUnixTimestamp {
		t.Errorf("Expected %v; got %v", expectedUnixTimestamp, *result)
	}
}

func TestToUnix_WhenTimeNil_ShouldReturnNil(t *testing.T) {
	// Arrange
	var nilTime *time.Time

	// Act
	result := utils.ToUnix(nilTime)

	// Assert
	if result != nil {
		t.Errorf("Expected %v; got %v", nil, result)
	}
}

func TestToTime_WhenValidUnixTimestamp_ShouldReturnValidTime(t *testing.T) {
	// Arrange
	validTimestamp := int64(1553811586000)

	// Act
	result := utils.ToTime(&validTimestamp)

	// Assert
	expectedTime := time.Date(2019, time.March, 28, 22, 19, 46, 0, time.FixedZone("UTC", 0))
	if !result.Equal(expectedTime) {
		t.Errorf("Expected %v; got %v", expectedTime, result)
	}
}

func TestToTime_WhenUnixTimestampNil_ShouldReturnNil(t *testing.T) {
	// Arrange
	var nilTimestamp *int64

	// Act
	result := utils.ToTime(nilTimestamp)

	// Assert
	if result != nil {
		t.Errorf("Expected %v; got %v", nil, result)
	}
}

func TestToUnixAndToTime_WhenConvertedToUnixAndBackToTime_ShouldReturnTheOriginalTime(t *testing.T) {
	// Arrange
	startingTime := time.Now().UTC().Round(time.Millisecond)

	// Act
	result := utils.ToTime(utils.ToUnix(&startingTime))

	// Assert
	if *result != startingTime {
		t.Errorf("Expected %v; got %v", startingTime, result)
	}
}
