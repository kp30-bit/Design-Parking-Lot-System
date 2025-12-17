# Parking Management System

A low-level design (LLD) implementation of a parking management system written in Go, demonstrating clean architecture principles, design patterns, and SOLID principles.

## Table of Contents

- [Architecture](#architecture)
- [Design Patterns](#design-patterns)
- [SOLID Principles](#solid-principles)
- [Project Structure](#project-structure)
- [How to Run](#how-to-run)
- [Usage Example](#usage-example)

## Architecture

This project follows a **clean code architecture** pattern, organizing code into distinct layers with clear responsibilities:

### Layer Structure

```
┌─────────────────────────────────────┐
│         Application Layer           │
│         (cmd/main.go)               │
└─────────────────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────┐
│         Managers Layer              │
│  (ParkingLotMgr, StrategyMgr,      │
│   TicketMgr)                        │
└─────────────────────────────────────┘
                  │
        ┌─────────┴─────────┐
        ▼                   ▼
┌──────────────┐    ┌─────────────────┐
│ Domain Layer │    │  Usecase Layer  │
│ (Entities)   │    │  (Strategies)   │
└──────────────┘    └─────────────────┘
        │                   │
        └─────────┬─────────┘
                  ▼
        ┌─────────────────┐
        │ Interface Layer │
        │  (Contracts)    │
        └─────────────────┘
```

### Layer Descriptions

1. **Domain Layer** (`internal/domain/`)
   - Contains core business entities: `ParkingLot`, `ParkingFloor`, `ParkingSlot`, `Vehicle`, `Ticket`
   - Represents the business logic and data structures
   - No dependencies on external layers

2. **Interface Layer** (`internal/interface/`)
   - Defines contracts/interfaces that decouple layers
   - `IParkingStrategy` interface enables strategy pattern implementation

3. **Managers Layer** (`internal/managers/`)
   - Orchestrates business operations
   - `ParkingLotMgr`: Main orchestrator for parking operations
   - `StrategyMgr`: Factory for creating parking strategies
   - `TicketMgr`: Manages ticket lifecycle

4. **Usecase Layer** (`internal/usecase/`)
   - Contains specific business use cases/strategies
   - `ClosestAvailableParkingStrategy`: Implements closest-first parking logic

5. **Application Layer** (`cmd/`)
   - Entry point of the application
   - Initializes and wires components together

## Design Patterns

### 1. Strategy Pattern

**Purpose**: Allows selecting parking algorithms at runtime.

**Implementation**:
- **Interface**: `IParkingStrategy` in `internal/interface/parking_strategy.go`
  ```go
  type IParkingStrategy interface {
      GetFreeSlot(slotType domain.SlotType) (*domain.ParkingSlot, error)
  }
  ```

- **Concrete Strategy**: `ClosestAvailableParkingStrategy` in `internal/usecase/closest_available_parking.go`
  - Implements the interface to find the closest available parking slot

- **Usage**: In `ParkingLotMgr.Park()`, the strategy is selected and used:
  ```go
  parkingStrategy := p.StrategyMgr.SelectStrategy(strategy, p.ParkingLot.GetParkingFloorMap())
  parkingSlot, err := parkingStrategy.GetFreeSlot(v.GetSlot())
  ```

**Benefits**:
- Easy to add new parking strategies (e.g., `RandomAvailableParking`, `CheapestParking`)
- Decouples parking logic from the main parking manager
- Follows Open/Closed Principle

### 2. Factory Pattern

**Purpose**: Creates appropriate strategy objects based on runtime input.

**Implementation**: `StrategyMgr.SelectStrategy()` in `internal/managers/strategy_mgr.go`
```go
func (sm *StrategyMgr) SelectStrategy(inputStrategy domain.ParkingStrategy, ParkingFloorMap map[int]*domain.ParkingFloor) interfaces.IParkingStrategy {
    switch inputStrategy {
    case domain.ClosestAvailableParking:
        return usecase.NewClosestAvailableParkingStrategy(ParkingFloorMap)
    }
    return nil
}
```

**Benefits**:
- Encapsulates object creation logic
- Centralizes strategy instantiation
- Makes it easy to add new strategies without modifying client code

### 3. Singleton Pattern

**Purpose**: Ensures only one instance of `ParkingLotMgr` exists throughout the application lifecycle.

**Implementation**: In `internal/managers/parking_lot_mgr.go`
```go
var (
    ParkingLotMgrInstance *ParkingLotMgr
    singletonInstance     sync.Once
)

func NewParkingLotMgr(pl *domain.ParkingLot, st *StrategyMgr, t *TicketMgr) *ParkingLotMgr {
    singletonInstance.Do(func() {
        ParkingLotMgrInstance = &ParkingLotMgr{
            ParkingLot:  pl,
            StrategyMgr: st,
            TicketMgr:   t,
        }
    })
    return ParkingLotMgrInstance
}
```

**Benefits**:
- Ensures single point of control for parking operations
- Thread-safe initialization using `sync.Once`
- Prevents multiple instances that could cause state inconsistencies

### 4. Interface Segregation Pattern

**Purpose**: Defines focused interfaces that clients depend on.

**Implementation**: `IParkingStrategy` interface contains only the method needed for parking slot selection:
```go
type IParkingStrategy interface {
    GetFreeSlot(slotType domain.SlotType) (*domain.ParkingSlot, error)
}
```

**Benefits**:
- Clients only depend on methods they actually use
- Easy to implement and test
- Reduces coupling between components

## SOLID Principles

### 1. Single Responsibility Principle (SRP)

Each class/struct has a single, well-defined responsibility:

- **`ParkingLot`**: Manages parking lot structure (floors, slots) and basic park/unpark operations
- **`ParkingLotMgr`**: Orchestrates parking operations, delegates to strategies and ticket management
- **`StrategyMgr`**: Responsible only for creating strategy instances
- **`TicketMgr`**: Manages ticket storage and retrieval
- **`ClosestAvailableParkingStrategy`**: Implements only the closest-first parking algorithm

**Example**: `ParkingLotMgr` doesn't handle ticket storage directly; it delegates to `TicketMgr`:
```go
func (p *ParkingLotMgr) Park(v domain.Vehicle, strategy domain.ParkingStrategy) (*domain.Ticket, error) {
    // ... parking logic ...
    p.TicketMgr.AddTicket(ticket)  // Delegates ticket management
    return ticket, nil
}
```

### 2. Open/Closed Principle (OCP)

The system is open for extension but closed for modification:

- **Adding new parking strategies**: Create a new struct implementing `IParkingStrategy` and add a case in `StrategyMgr.SelectStrategy()` without modifying existing code
- **Adding new vehicle types**: Implement the `Vehicle` interface without changing existing vehicle handling code

**Example**: To add a `RandomAvailableParking` strategy:
1. Create `RandomAvailableParkingStrategy` implementing `IParkingStrategy`
2. Add case in `StrategyMgr.SelectStrategy()`
3. No changes needed in `ParkingLotMgr` or other components

### 3. Liskov Substitution Principle (LSP)

Subtypes must be substitutable for their base types:

- **Vehicle Interface**: `Car` and `Bike` can be used interchangeably wherever `Vehicle` is expected
- **Strategy Interface**: Any implementation of `IParkingStrategy` can replace another without breaking functionality

**Example**: Both `Car` and `Bike` implement `Vehicle` interface:
```go
type Vehicle interface {
    GetNo() int
    GetType() VehicleType
    GetSlot() SlotType
}
```

Any method accepting `Vehicle` can work with both `Car` and `Bike` instances.

### 4. Interface Segregation Principle (ISP)

Clients should not be forced to depend on interfaces they don't use:

- **`IParkingStrategy`**: Contains only `GetFreeSlot()` method, which is all that's needed
- **`Vehicle`**: Contains only methods relevant to vehicle operations

**Example**: `IParkingStrategy` is minimal and focused:
```go
type IParkingStrategy interface {
    GetFreeSlot(slotType domain.SlotType) (*domain.ParkingSlot, error)
}
```

No unnecessary methods that implementations would be forced to provide.

### 5. Dependency Inversion Principle (DIP)

High-level modules should not depend on low-level modules; both should depend on abstractions:

- **`ParkingLotMgr`** depends on `IParkingStrategy` interface, not concrete strategy implementations
- **Strategy creation** happens through `StrategyMgr`, which returns interface types
- **Domain entities** don't depend on managers or usecases

**Example**: `ParkingLotMgr` uses interface abstraction:
```go
func (p *ParkingLotMgr) Park(v domain.Vehicle, strategy domain.ParkingStrategy) (*domain.Ticket, error) {
    parkingStrategy := p.StrategyMgr.SelectStrategy(strategy, ...)  // Returns IParkingStrategy interface
    parkingSlot, err := parkingStrategy.GetFreeSlot(v.GetSlot())     // Uses interface method
    // ...
}
```

The manager doesn't know about `ClosestAvailableParkingStrategy` specifically; it only knows about the interface.

## Project Structure

```
.
├── cmd/
│   └── main.go                    # Application entry point
├── internal/
│   ├── domain/                    # Domain entities
│   │   ├── parking_lot.go         # ParkingLot entity
│   │   ├── parking_floor.go       # ParkingFloor entity
│   │   ├── parking_slot.go        # ParkingSlot entity
│   │   ├── vehicle.go             # Vehicle interface and implementations
│   │   ├── ticket.go              # Ticket entity
│   │   └── util.go                # Utility types and constants
│   ├── interface/                 # Interface definitions
│   │   └── parking_strategy.go    # IParkingStrategy interface
│   ├── managers/                  # Business logic managers
│   │   ├── parking_lot_mgr.go     # Main parking orchestrator
│   │   ├── strategy_mgr.go        # Strategy factory
│   │   └── ticket_mgr.go          # Ticket manager
│   └── usecase/                   # Use case implementations
│       └── closest_available_parking.go  # Closest parking strategy
├── .vscode/
│   └── launch.json                # VS Code debug configuration
├── go.mod                         # Go module definition
└── README.md                      # This file
```

## How to Run

### Prerequisites

- Go 1.25.2 or later

### Running the Application

1. **Clone the repository** (if applicable)

2. **Run the application**:
   ```bash
   go run cmd/main.go
   ```

### Debugging with VS Code

1. Open the project in VS Code
2. Go to Run and Debug (F5)
3. Select "Debug Parking Mgmt" configuration
4. Set breakpoints and debug

## Usage Example

```go
// Create vehicles
car := domain.NewCar(123, domain.CarType)
bike := domain.NewBike(456, domain.BikeType)

// Create parking slots
parkingSlot1 := domain.NewParkingSlot(1, domain.BikeSlot, nil, true)
parkingSlot2 := domain.NewParkingSlot(1, domain.CarSlot, nil, true)

// Create parking floors
level1 := domain.NewParkingFloor(1)
level1.AddParkingSlot(parkingSlot1)
level1.AddParkingSlot(parkingSlot2)

// Create parking lot
parkingLot := domain.NewParkingLot()
parkingLot.AddFloor(level1)

// Initialize managers
ticketMgr := managers.NewTicketMgr()
strategyMgr := managers.NewStrategyMgr()
parkingLotMgr := managers.NewParkingLotMgr(parkingLot, strategyMgr, ticketMgr)

// Park a vehicle
ticket, err := parkingLotMgr.Park(car, domain.ClosestAvailableParking)

// Unpark a vehicle
parkingLotMgr.Unpark(ticket)

// Show all parked vehicles
parkingLotMgr.ParkingLot.ShowAllParkedVehicles()
```

## Key Features

- ✅ Layered architecture for maintainability
- ✅ Strategy pattern for flexible parking algorithms
- ✅ Factory pattern for object creation
- ✅ Singleton pattern for resource management
- ✅ SOLID principles throughout
- ✅ Interface-based design for extensibility
- ✅ Thread-safe singleton implementation

## Future Enhancements

- Add more parking strategies (Random, Cheapest, etc.)
- Implement pricing strategies
- Add support for more vehicle types
- Implement payment processing
- Add database persistence
- Add REST API layer
- Implement comprehensive unit tests
