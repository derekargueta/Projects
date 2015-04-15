
/**
 * @author Derek Argueta
 */
public class Main {

    public static void main(String[] args) {
        if(args.length == 0) {
            /**
             * TODO - default runs GUI
             */
        } else {
            if(args.length != 3) {
                System.out.println("Incorrect arguments.\n\nUsage:\n<url> <integer duration> <seconds|minutes|hours>");
                System.exit(-1);
            }
            String urlString = args[0];
            if(!urlString.startsWith("http://") || !urlString.startsWith("https://")) {
                System.out.println("Method not found, applying http://");
                urlString = "http://" + urlString;
            }

            String type = args[2];
            try {
                int pause = Integer.parseInt(args[1]);
                if(type.equals("seconds")) {
                    pause *= 1000;
                } else if(type.equals("minutes")) {
                    pause *= 60000;
                } else if(type.equals("hours")) {
                    pause *= 360000;
                } else {
                    System.out.println("Unknown time type! Please specify seconds, minutes, or hours");
                    System.exit(-1);
                }

                while(true) {
                    new ConnectThread(urlString).start();
                    try {
                        Thread.sleep(pause);
                    } catch (InterruptedException e) {
                        System.out.println("Program has been interrupted. Goodbye!");
                    }
                }
            } catch(NumberFormatException e) {
                e.printStackTrace();
                System.exit(-1);
            }
        }
    }
}
