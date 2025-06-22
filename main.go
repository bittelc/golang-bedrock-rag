package main

import (
	"context"
	"log"

	"golang-bedrock-rag/aws"
)

// Configuration placeholders
const (
	EmbeddingModelID = "amazon.titan-embed-text-v1"
	ChatModelID      = "anthropic.claude-v2"

	PineconeEndpoint = "https://YOUR-PINECONE-PROJECT.svc.YOUR-REGION.pinecone.io"
	PineconeKey      = "YOUR_PINECONE_API_KEY"
	IndexName        = "documents-index"
)

func main() {
	ctx := context.Background()
	_, err := aws.AuthToAws(&ctx)
	if err != nil {
		log.Fatalf("could not auth to AWS:", err)
	}

	// bedrockClient := bedrockruntime.NewFromConfig(cfg)
	// spew.Dump(bedrockClient)

	// 	docsDir := "./docs"
	// 	files, err := filepath.Glob(filepath.Join(docsDir, "*.docx"))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	for _, file := range files {
	// 		text, err := extractTextFromDocx(file)
	// 		if err != nil {
	// 			log.Printf("failed to extract %s: %v", file, err)
	// 			continue
	// 		}
	// 		chunks := chunkText(text, 500)
	// 		for i, chunk := range chunks {
	// 			vec, err := embedText(ctx, bedrockClient, chunk)
	// 			if err != nil {
	// 				log.Printf("embedding error: %v", err)
	// 				continue
	// 			}
	// 			metadata := map[string]string{"source": filepath.Base(file), "chunk": fmt.Sprint(i)}
	// 			if err := upsertToPinecone(chunkID(file, i), vec, metadata); err != nil {
	// 				log.Printf("pinecone upsert error: %v", err)
	// 			}
	// 		}
	// 	}

	// 	reader := bufio.NewReader(os.Stdin)
	// 	for {
	// 		fmt.Print("\nYou: ")
	// 		question, _ := reader.ReadString('\n')
	// 		question = strings.TrimSpace(question)
	// 		if question == "exit" {
	// 			break
	// 		}

	// 		topChunks, err := queryPinecone(question, 3)
	// 		if err != nil {
	// 			log.Printf("retrieval error: %v", err)
	// 			continue
	// 		}
	// 		contextText := strings.Join(topChunks, "\n---\n")
	// 		answer, err := chatWithClaude(ctx, bedrockClient, question, contextText)
	// 		if err != nil {
	// 			log.Printf("chat error: %v", err)
	// 			continue
	// 		}
	// 		fmt.Printf("Assistant: %s\n", answer)
	// 	}
	// }

	// func extractTextFromDocx(path string) (string, error) {
	// 	doc, err := document.Open(path)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	defer doc.Close()

	// 	var sb strings.Builder
	// 	for _, para := range doc.Paragraphs() {
	// 		sb.WriteString(para.Text())
	// 		sb.WriteString("\n")
	// 	}
	// 	return sb.String(), nil
	// }

	// func chunkText(text string, maxLen int) []string {
	// 	words := strings.Fields(text)
	// 	var chunks []string
	// 	for i := 0; i < len(words); i += maxLen {
	// 		end := i + maxLen
	// 		if end > len(words) {
	// 			end = len(words)
	// 		}
	// 		chunks = append(chunks, strings.Join(words[i:end], " "))
	// 	}
	// 	return chunks
	// }

	// func embedText(ctx context.Context, client *bedrockruntime.Client, text string) ([]float32, error) {
	// 	input := map[string]interface{}{"inputText": text}
	// 	body, _ := json.Marshal(input)

	// 	resp, err := client.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
	// 		ModelId:     aws.String(EmbeddingModelID),
	// 		ContentType: aws.String("application/json"),
	// 		Body:        bytes.NewReader(body),
	// 	})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	defer resp.Body.Close()

	// 	var out struct {
	// 		Embedding []float32 `json:"embedding"`
	// 	}
	// 	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
	// 		return nil, err
	// 	}
	// 	return out.Embedding, nil
	// }

	// func chatWithClaude(ctx context.Context, client *bedrockruntime.Client, question, contextText string) (string, error) {
	// 	prompt := fmt.Sprintf("Human: Based on the context below, answer the question.\n\nContext:\n%s\nQuestion:\n%s\nAssistant:", contextText, question)
	// 	input := map[string]interface{}{
	// 		"prompt":               prompt,
	// 		"max_tokens_to_sample": 300,
	// 		"temperature":          0.7,
	// 		"stop_sequences":       []string{"\n\nHuman:"},
	// 	}
	// 	body, _ := json.Marshal(input)

	// 	resp, err := client.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
	// 		ModelId:     aws.String(ChatModelID),
	// 		ContentType: aws.String("application/json"),
	// 		Body:        bytes.NewReader(body),
	// 	})
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	defer resp.Body.Close()

	// 	var out struct {
	// 		Completion string `json:"completion"`
	// 	}
	// 	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
	// 		return "", err
	// 	}
	// 	return out.Completion, nil
	// }

	// func chunkID(file string, idx int) string {
	// 	base := filepath.Base(file)
	// 	return fmt.Sprintf("%s-%d", base, idx)
	// }

	// func upsertToPinecone(id string, vector []float32, metadata map[string]string) error {
	// 	req := map[string]interface{}{
	// 		"vectors":   []map[string]interface{}{{"id": id, "values": vector, "metadata": metadata}},
	// 		"namespace": IndexName,
	// 	}
	// 	body, _ := json.Marshal(req)
	// 	reqURL := fmt.Sprintf("%s/vectors/upsert", PineconeEndpoint)
	// 	request, _ := http.NewRequest("POST", reqURL, bytes.NewReader(body))
	// 	request.Header.Set("Api-Key", PineconeKey)
	// 	request.Header.Set("Content-Type", "application/json")

	// 	resp, err := http.DefaultClient.Do(request)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	io.Copy(io.Discard, resp.Body)
	// 	resp.Body.Close()
	// 	return nil
	// }

	// func queryPinecone(query string, topK int) ([]string, error) {
	// 	vec, err := embedText(context.Background(), bedrockruntime.NewFromConfig(aws.Config{}), query)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	reqBody := map[string]interface{}{
	// 		"vector":    vec,
	// 		"topK":      topK,
	// 		"namespace": IndexName,
	// 	}
	// 	body, _ := json.Marshal(reqBody)
	// 	reqURL := fmt.Sprintf("%s/query", PineconeEndpoint)
	// 	request, _ := http.NewRequest("POST", reqURL, bytes.NewReader(body))
	// 	request.Header.Set("Api-Key", PineconeKey)
	// 	request.Header.Set("Content-Type", "application/json")

	// 	resp, err := http.DefaultClient.Do(request)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	defer resp.Body.Close()

	//	var out struct {
	//		Matches []struct {
	//			Score    float64           `json:"score"`
	//			ID       string            `json:"id"`
	//			Metadata map[string]string `json:"metadata"`
	//		} `json:"matches"`
	//	}
	//
	//	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
	//		return nil, err
	//	}
	//
	// var texts []string
	//
	//	for _, m := range out.Matches {
	//		texts = append(texts, fmt.Sprintf("[Source: %s | Chunk: %s]", m.Metadata["source"], m.Metadata["chunk"]))
	//	}
	//
	// return texts, nil
}
