package main

import (
	"testing"

	"github.com/aruncs31s/pdf"
)

func TestBoldDetection(t *testing.T) {
	f, r, err := pdf.Open("test/dms.pdf")
	if err != nil {
		t.Fatalf("Failed to open test PDF: %v", err)
	}
	defer f.Close()

	if r.NumPage() == 0 {
		t.Fatal("Expected at least 1 page")
	}

	p := r.Page(1)
	fontNames := p.Fonts()
	if len(fontNames) == 0 {
		t.Fatal("Expected at least 1 font on page 1")
	}

	for _, name := range fontNames {
		_ = p.Font(name).IsBold()
	}

	texts, err := r.GetStyledTexts()
	if err != nil {
		t.Fatalf("GetStyledTexts failed: %v", err)
	}
	if len(texts) == 0 {
		t.Fatal("GetStyledTexts returned no text")
	}

	for _, txt := range texts {
		if txt.S == "" {
			t.Error("Expected non-empty S in Text")
		}
	}
}

func TestFontIsBold(t *testing.T) {
	f, r, err := pdf.Open("test/dms.pdf")
	if err != nil {
		t.Fatalf("Failed to open test PDF: %v", err)
	}
	defer f.Close()

	p := r.Page(1)
	for _, name := range p.Fonts() {
		font := p.Font(name)
		isBold := font.IsBold()
		t.Logf("Font %s: BaseFont=%q, IsBold=%v", name, font.BaseFont(), isBold)
	}
}

func TestTextIsBoldField(t *testing.T) {
	f, r, err := pdf.Open("test/dms.pdf")
	if err != nil {
		t.Fatalf("Failed to open test PDF: %v", err)
	}
	defer f.Close()

	texts, err := r.GetStyledTexts()
	if err != nil {
		t.Fatalf("GetStyledTexts failed: %v", err)
	}

	nBold := countBold(texts)
	nNonBold := len(texts) - nBold
	if nNonBold == 0 {
		t.Error("Expected at least some non-bold text")
	}

	t.Logf("Text segments: %d total, %d bold, %d non-bold",
		len(texts), nBold, nNonBold)
}

func countBold(texts []pdf.Text) int {
	n := 0
	for _, t := range texts {
		if t.IsBold {
			n++
		}
	}
	return n
}

func TestContentIsBold(t *testing.T) {
	f, r, err := pdf.Open("test/dms.pdf")
	if err != nil {
		t.Fatalf("Failed to open test PDF: %v", err)
	}
	defer f.Close()

	p := r.Page(1)
	content := p.Content()
	if len(content.Text) == 0 {
		t.Fatal("Expected text content on page 1")
	}

	seen := 0
	for _, txt := range content.Text {
		if seen >= 10 {
			break
		}
		if txt.S == "" || txt.S == " " || txt.S == "\n" {
			continue
		}
		t.Logf("Font=%s FontSize=%.1f IsBold=%v S=%q",
			txt.Font, txt.FontSize, txt.IsBold, txt.S)
		seen++
	}
}
