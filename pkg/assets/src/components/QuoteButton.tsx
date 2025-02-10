import React, { useState } from "react"
import { Button } from "@/components/ui/button"
import { getQuote } from "../actions/getQuote"

export default function QuoteButton() {
  const [quote, setQuote] = useState<{ text: string; author: string } | null>(null)
  const [isLoading, setIsLoading] = useState(false)

  const handleClick = async () => {
    setIsLoading(true)
    try {
      const newQuote = await getQuote()
      setQuote(newQuote)
    } catch (error) {
      console.error("Failed to fetch quote:", error)
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="flex flex-col items-center gap-4">
      <Button onClick={handleClick} disabled={isLoading}>
        {isLoading ? "Fetching Quote..." : "Get Quote"}
      </Button>
      {quote && (
        <div className="max-w-md text-center">
          <p className="text-lg font-semibold">{quote.text}</p>
          <p className="text-sm text-gray-500">- {quote.author}</p>
        </div>
      )}
    </div>
  )
}
