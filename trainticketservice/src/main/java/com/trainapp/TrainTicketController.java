package com.trainapp;

import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;

@RestController
@RequestMapping("/tickets")
public class TrainTicketController {

    private List<Receipt> receipts = new ArrayList<>();
    private AtomicInteger seatCounter = new AtomicInteger(1);

    @PostMapping("/purchase")
public Receipt purchaseTicket(@RequestBody User user, @RequestParam String from, @RequestParam String to, @RequestParam double price) {
    String seat = allocateSeat();
    user.setSeat(seat);
    Receipt receipt = new Receipt(from, to, user, price, seat);
    receipts.add(receipt);
    return receipt;
}
    @GetMapping("/receipt/{email}")
    public Receipt getReceipt(@PathVariable String email) {
        return receipts.stream()
                .filter(receipt -> receipt.getUser().getEmail().equals(email))
                .findFirst()
                .orElse(null);
    }

    @GetMapping("/users/{section}")
    public List<User> getUsersBySection(@PathVariable String section) {
        List<User> usersInSection = new ArrayList<>();
        for (Receipt receipt : receipts) {
            if (receipt.getUser().getSeat().startsWith(section)) {
                usersInSection.add(receipt.getUser());
            }
        }
        return usersInSection;
    }

    @DeleteMapping("/remove/{email}")
    public void removeUser(@PathVariable String email) {
        receipts.removeIf(receipt -> receipt.getUser().getEmail().equals(email));
    }

    @PutMapping("/modify-seat")
public void modifySeat(@RequestParam String email, @RequestParam String newSeat) {
    for (Receipt receipt : receipts) {
        if (receipt.getUser().getEmail().equals(email)) {
            System.out.println("Found user: " + email);
            receipt.getUser().setSeat(newSeat);
            receipt.setSeat(newSeat);
            System.out.println("Changed seat to: " + newSeat);
            break;
        }
    }
}

    private String allocateSeat() {
        int seatNumber = seatCounter.getAndIncrement();
        if (seatNumber % 2 == 0) {
            return "B" + seatNumber;
        } else {
            return "A" + seatNumber;
        }
    }
}
