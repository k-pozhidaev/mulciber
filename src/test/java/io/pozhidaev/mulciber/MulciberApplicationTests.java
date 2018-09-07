package io.pozhidaev.mulciber;

import io.pozhidaev.mulciber.config.SecurityConfig;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit4.SpringRunner;

import javax.sql.DataSource;

@RunWith(SpringRunner.class)
@SpringBootTest
public class MulciberApplicationTests {

	@MockBean
	protected DataSource dataSource;

	@MockBean
	protected SecurityConfig securityConfig;

	@Test
	public void contextLoads() {
	}

}
