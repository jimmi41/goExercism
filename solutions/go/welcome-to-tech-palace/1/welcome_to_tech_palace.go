package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    res := ""

    for i := 0; i < numStarsPerLine; i++ {
        res += "*"
    }
    res += "\n"

    res += welcomeMsg + "\n"

    for i := 0; i < numStarsPerLine; i++ {
        res += "*"
    }

    return res
}



func CleanupMessage(oldMsg string) string {
    // Remove stars
    msg := strings.ReplaceAll(oldMsg, "*", "")

    // Trim leading/trailing whitespace
    msg = strings.TrimSpace(msg)

    return msg
}