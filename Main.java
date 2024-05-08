import java.util.Stack;

public class Main {
    public static void main(String[] args) {
       Stack<String> stack = new Stack<String>();

       System.out.println(stack.empty());

       stack.push("Minecraft");
       stack.push("Call of duty");
       stack.push("Blade");
       stack.push("Battlefield 3");

       System.out.println(stack);

       String myFavGame = stack.pop();

       System.out.println(myFavGame);


       System.out.println(stack.peek());

       System.out.println(stack.search("Minecraft"));

       



    }

}


