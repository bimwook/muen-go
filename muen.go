package muen;
import(
  "io"
  "os"
  "net"
  "time"
  "bufio"
  "strings"
  "strconv"
  "math/rand"
);

type MuenError struct{
  Code int
  Info string
}

func (this MuenError) Error() string{
  return "[" + strconv.Itoa(this.Code) + "] " + this.Info;
}

//COMMON FUNCTIONS
func Now() string {
  return time.Now().Format("2006-01-02 15:04:05");
}

func Rndid() string {
  return time.Now().Format("20060102150405");
}

func Send(str string) bool {
  socket, err := net.Dial("udp4", "127.0.0.1:19800");
  if err != nil {
    return false;
  }
  io.WriteString(socket, str);
  defer socket.Close();
  return true;
}

func Sendln(str string) bool {
  socket, err := net.Dial("udp4", "127.0.0.1:19800");
  if err != nil {
    return false;
  }
  io.WriteString(socket, "[" + Now() + "] " + str + "\r\n");
  defer socket.Close();
  return true;
}

func NewKey() string{
  ret := time.Now().Format("20060102");
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  for i:=0; i<40; i++ {
    ret = ret + strconv.Itoa(r.Intn(10));
  }
  return ret;
}

func SubString(str string,begin,length int) string {  
  s := []rune(str);  
  size:= len(s);  
  if begin < 0 {  
    begin = 0;  
  }  
  if begin >= size {  
    begin = size;  
  }  
  end := begin + length;  
  if end > size {  
    end = size;
  }  
  return string(s[begin:end]);  
}

func LoadMap(fn string) (ret map[string]string, ok bool) {
  ok = true;
  file, e := os.Open(fn) // For read access.
  if e != nil {
    ok = false;
    return nil, ok;
  }
  ret= make(map[string]string);
  reader := bufio.NewReader(file);
  for{
    r, e := reader.ReadString('\n');
    line := strings.Replace(strings.Replace(r, "\r", "", -1), "\n", "", -1);
    if e!=nil {
      break;
    }
    p:= strings.Index(line, "=");
    if(p!=-1){
      key := SubString(line, 0, p);
      value:= SubString(line, p + 1, len(line));
      ret[key] = value;
    }
  }
  defer file.Close();
  
  return ret, ok;
}

func HtmlEncode(s string) string{
  ret:= strings.Replace(s, "&" , "&amp;", -1);
  ret = strings.Replace(ret, "<", "&lt;", -1);
  ret = strings.Replace(ret, ">", "&gt;", -1);
  return ret;
}

//TYPE ROOT
type Root struct{
  Name string
  Version string
  Tag int
};

func (this Root) Now() string {
  return time.Now().Format("2006-01-02 15:04:05");
}

func (this Root) Log(str string) string {
  return "[" + this.Now() + "] " + this.Name + ": " + str;
}

func (this *Root) SetName(name string) bool {
  this.Name = name;
  return true;
} 

func (this *Root) SetVersion (version string) bool{
  this.Version = version;
  return true;
}
