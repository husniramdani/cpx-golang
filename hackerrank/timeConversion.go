package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

 // basic logic
func timeConversion2(s string) string {
  // Write your code here
  // am - pagi - siang
  // pm - siang - malam
  // 12.01.00PM => 12.01.00
  // 12:01.00AM => 00.01.00
  // if 12 && PM => hapus PM
  // if 12 && AM => -12 / 00
  // if AM => hapus AM
  // if PM => +12 => 11PM => 23
  result := s
  hours := s[:2]
  AM := strings.Contains(s, "AM")
  PM := strings.Contains(s, "PM")
  isTwelve := strings.Contains(s, "12")

  // delete the time format
  if AM {
    if isTwelve {
      result = strings.Replace(result, hours, "00", -1)
    }
    result = strings.Replace(result, "AM", "", -1)
  }
  if PM {
    if !isTwelve {
      count, _ := strconv.Atoi(hours)
      reformat := strconv.Itoa(count+12)
      result = strings.Replace(result, hours, reformat, -1)
    }
    result = strings.Replace(result, "PM", "", -1)
  }
  return result
}

// better logic
func timeConversion(s string) string {
  hours, _ := strconv.Atoi(s[:2])
  isAM := s[8:] == "AM"

  if hours == 12 && isAM {
    return "00" + s[2:8]
  }
  if hours != 12 && !isAM {
    return fmt.Sprintf("%d%s", hours + 12, s[2:8])
  }
  return s[:8]
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    // stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    // checkError(err)

    // defer stdout.Close()

    // writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Println(result)
    // fmt.Fprintf(writer, "%s\n", result)

    // writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

