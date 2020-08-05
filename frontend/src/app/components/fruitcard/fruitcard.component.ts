import { CartComponent } from './../cart/cart.component';
import { Customer } from './../../models/customer.model';
import { Component, OnInit, Input } from '@angular/core';
import { FruitModel } from 'src/app/models/fruit.model';
import { CartService } from './../../services/cart.service';
import { AuthenticationService } from './../../services/authentication.service';
import { DiscountService } from './../../services/discount.service';
import { FruitService } from 'src/app/services/fruit.service';
import { Observable, pipe } from 'rxjs';
import { Datum } from './../../models/cartitem.model';
import { map } from 'rxjs/operators';


@Component({
  selector: 'app-fruitcard',
  templateUrl: './fruitcard.component.html',
  styleUrls: ['./fruitcard.component.scss']
})
export class FruitcardComponent implements OnInit {

  @Input()
  fruit: FruitModel
  currentUser: Customer;
  cartList: Observable<Datum[]>;

  count: number;

  constructor(private cartService: CartService, 
    public authenticationService: AuthenticationService, 
    public discountService: DiscountService,
    public fruitService: FruitService,
    ) {
    this.authenticationService.currentUser.subscribe(x => {
      this.currentUser = x
    });

  }

  

  ngOnInit(): void {
    this.count = 0;
  }

  addToCount(): void {
    this.count++;
  }

  subToCount(): void {
    this.count--;
  }

  addToCart(): void {
    this.cartService.addToCart(this.currentUser.Cart.ID, this.fruit, this.count).subscribe(() => {
      
      this.cartService.update.next(true)
      this.discountService.update.next(true)
    })
  }

}
