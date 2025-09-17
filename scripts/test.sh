#!/bin/bash
set -e

echo "🧪 Running tests for Feast Friends API..."

# Run unit tests with coverage
echo "🔬 Running unit tests..."
go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

# Generate coverage report
echo "📊 Generating coverage report..."
go tool cover -html=coverage.out -o coverage.html

# Check coverage threshold (80%)
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
THRESHOLD=80

echo "📈 Test coverage: ${COVERAGE}%"

if (( $(echo "$COVERAGE >= $THRESHOLD" | bc -l) )); then
    echo "✅ Coverage meets threshold (${THRESHOLD}%)"
else
    echo "❌ Coverage below threshold (${THRESHOLD}%): ${COVERAGE}%"
    exit 1
fi

# Run integration tests if DATABASE_URL is set
if [ ! -z "$DATABASE_URL" ]; then
    echo "🔗 Running integration tests..."
    go test -v -tags=integration ./tests/integration/...
fi

echo "✅ All tests passed!"