package arquivos;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.attribute.FileTime;
import java.nio.file.attribute.UserPrincipal;

public class Arquivos {
    
public static void main(String[] args) throws IOException, URISyntaxException {


    Path path = Path.of("c:\\Users\\henri\\dev\\pessoal\\java\\arquivos\\exemplo1.txt");
    System.out.println(path);
    FileTime lastModifiedTime = Files.getLastModifiedTime(path);
    System.out.println("lastModifiedTime = " + lastModifiedTime);
    UserPrincipal owner = Files.getOwner(path);
    System.out.println("owner = " + owner);

}

}