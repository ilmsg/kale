# kale

1. copy ไพล์ .env มาจากไพล์ตัวอย่าง

    $ cp .env.example .env

2. เพิ่ม `GEMINI_API_KEY` ในไพล์ `.env`

refer: [GenerateContent-TextOnly](https://pkg.go.dev/github.com/google/generative-ai-go/genai#example-GenerativeModel.GenerateContent-TextOnly)

---
ตัวอย่างใน source code

    // จัดเตรียม context สำหรับใช้สร้าง gemini client
    ctx := context.Background()

    // รับค่า GEMINI_API_KEY จาก environment มาใส่ option สำหรับใช้สร้าง gemini client
	opts := option.WithAPIKey(os.Getenv("GEMINI_API_KEY"))

    // สร้าง gemini client จาก context, option ที่เตรียมไว้แล้วก่อนหน้านี้
	client, err := genai.NewClient(ctx, opts)
	
    // จัดการกับ(ข้อผิดพลาด) error ถ้าหากเกิดขึ้น ในการสร้าง gemini client
    if err != nil {
		log.Fatal(err)
	}

    // ปิด gemini client ในตอนจบ function
	defer client.Close()

    // กำหนดชื่อ model ที่ต้องการใช้
	model := client.GenerativeModel("gemini-1.5-flash")

    // เรียกใช้การ Generate Content ด้วยคำถาม "Why be shy"
	resp, err := model.GenerateContent(ctx, genai.Text("Why be shy"))
	
    // จัดการกับ(ข้อผิดพลาด) error ถ้าหากเกิดขึ้น ในการ Generate Content
    if err != nil {
		log.Fatal(err)
	}

    // เรียกใช้ function แสดงผล
	printGenerateContentResponse(resp)

function อ่านค่าจาก .env มาใส่ environment

    func init() {
	    if err := godotenv.Load(); err != nil {
		    log.Fatal(err)
	    }
    }

function แสดงผลลัพธ์

    func printGenerateContentResponse(resp *genai.GenerateContentResponse) {
        for _, cand := range resp.Candidates {
            if cand.Content != nil {
                for _, part := range cand.Content.Parts {
                    fmt.Println(part)
                }
            }
        }
        println("---")
    }

---
**Why be shy**

Shyness is a complex emotion with a variety of underlying causes.  There's no single answer to "why be shy," but here are some contributing factors:

* **Genetics:** Some people are simply born with a temperament that predisposes them to shyness.  This can manifest as a heightened sensitivity to social situations and a greater tendency towards anxiety.

* **Temperament:**  Even without a genetic predisposition, a child's innate temperament can influence their social development.  A naturally more cautious or reserved child might develop shyness more easily than an outgoing one.

* **Early Childhood Experiences:**  Negative experiences in social situations, such as bullying, teasing, or rejection, can significantly impact a child's confidence and lead to shyness.  Lack of positive social interaction and support can also contribute.

* **Learned Behavior:**  Children might learn shyness by observing and imitating shy role models, such as parents or siblings.  They might also learn to associate social situations with negative consequences, reinforcing shy behavior.

* **Social Anxiety Disorder:**  In some cases, shyness can be a symptom of a more serious condition like social anxiety disorder.  This involves excessive fear and anxiety in social situations, significantly impacting daily life.

* **Low Self-Esteem:** Individuals with low self-esteem often lack confidence in their social abilities, making them more likely to feel shy and avoid social interactions.

* **Perfectionism:** The desire to present a perfect image can lead to anxiety and self-consciousness, making it difficult to relax and be oneself in social situations.

* **Neurobiological Factors:** Research suggests that certain brain regions and neurotransmitters play a role in shyness and social anxiety.

It's important to remember that shyness is a spectrum.  While some degree of shyness is normal, excessive shyness can significantly impact an individual's quality of life.  If shyness is causing distress or interfering with daily functioning, seeking professional help is advisable.
