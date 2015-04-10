import java.net.HttpURLConnection;
import java.net.URL;
import java.util.Date;

/**
 * @author Derek Argueta
 */
public class ConnectThread extends Thread {

    private String urlString;

    public ConnectThread(String url) {
        super();
        this.urlString = url;
    }

    public void run() {
        try {
            URL url = new URL(urlString);
            HttpURLConnection connection = (HttpURLConnection)url.openConnection();
            connection.setRequestMethod("GET");
            System.out.println("Attempting conection at " + new Date().toString());
            connection.connect();
            if(connection.getResponseCode() > 199 && connection.getResponseCode() < 300) {
                System.out.println("Successfully connected! Status: " + connection.getResponseCode());
            } else {
                System.out.println("uh oh, got a bad status code :/");
            }
        } catch (Exception e) {
            System.out.print("Network error");
        }
    }

}
